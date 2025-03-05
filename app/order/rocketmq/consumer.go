package rocketmq

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	"github.com/MakiJOJO/douyin-mall-echo/common/utils/config"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

// MessageHandler 定义消息处理函数类型
type MessageHandler func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error)

// ConsumerConfig 定义消费者配置
type ConsumerConfig struct {
	config.ConsumerConfig
	Handler MessageHandler
}

// InitConsumer 初始化通用的 RocketMQ 消费者
func InitConsumer(ctx context.Context, config ConsumerConfig) error {
	var c rocketmq.PushConsumer
	var err error

	// 创建消费者并进行重试
	for i := 0; i <= config.RetryCount; i++ {
		c, err = rocketmq.NewPushConsumer(
			consumer.WithGroupName(config.GroupName),
			consumer.WithNameServer(config.NsAddrs),
		)
		if err == nil {
			break
		}
		if i < config.RetryCount {
			mtl.Logger.Error(fmt.Sprintf("Failed to create new pull consumer (attempt %d/%d): %v, retrying in %v...", i+1, config.RetryCount+1, err, config.RetryTimeout))
			time.Sleep(config.RetryTimeout)
		}
	}
	if err != nil {
		mtl.Logger.Error(fmt.Sprintf("Failed to create new pull consumer after %d attempts", config.RetryCount+1), "retryErr", err)
		return err
	}

	// 订阅主题并进行重试
	for i := 0; i <= config.RetryCount; i++ {
		err = c.Subscribe(config.Topic, consumer.MessageSelector{}, config.Handler)
		if err == nil {
			break
		}
		if i < config.RetryCount {
			mtl.Logger.Warn("Failed to subscribe to topic %s (attempt %d/%d): %v, retrying in %v...", config.Topic, i+1, config.RetryCount+1, err, config.RetryTimeout)
			time.Sleep(config.RetryTimeout)
		}
	}
	if err != nil {
		mtl.Logger.Error("Failed to subscribe to topic %s after %d attempts: %v", config.Topic, config.RetryCount+1, err)
		return err
	}

	// 启动消费者并进行重试
	for i := 0; i <= config.RetryCount; i++ {
		err = c.Start()
		if err == nil {
			break
		}
		if i < config.RetryCount {
			mtl.Logger.Error(fmt.Sprintf("Failed to start consumer (attempt %d/%d): %v, retrying in %v...", i+1, config.RetryCount+1, err, config.RetryTimeout))
			time.Sleep(config.RetryTimeout)
		}
	}
	if err != nil {
		mtl.Logger.Error(fmt.Sprintf("Failed to start consumer after %d attempts: %v", config.RetryCount+1, err))
		return err
	}

	mtl.Logger.Info("Consumer %s subscribed to topic %s and started successfully", config.GroupName, config.Topic)

	// 监听系统信号，实现优雅关闭
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// 监听上下文取消信号
	go func() {
		select {
		case sig := <-signals:
			mtl.Logger.Info("Received signal %v, shutting down consumer...", sig)
		case <-ctx.Done():
			mtl.Logger.Info("Context cancelled, shutting down consumer...")
		}
		err := c.Shutdown()
		if err != nil {
			mtl.Logger.Error("Failed to shut down consumer: %v", err)
		} else {
			mtl.Logger.Error("Consumer shut down successfully")
		}
		os.Exit(0)
	}()

	return nil
}
