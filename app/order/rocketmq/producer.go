package rocketmq

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MakiJOJO/douyin-mall-echo/app/order/config"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var p rocketmq.Producer
var producerInitialized bool

func InitProducer() error {
	var err error
	p, err = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(config.GlobalConfig.RocketMq.NsAddrs)),
		producer.WithRetry(2),
		producer.WithGroupName("dymall"),
	)
	if err != nil {
		fmt.Printf("create producer error: %v", err)
		return err
	}

	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %v", err)
		return err
	}
	producerInitialized = true
	fmt.Println("Producer started successfully")
	return nil
}

func SendDelayedMessage(topic, message string, delay time.Duration) error {
	if !producerInitialized {
		mtl.Logger.Error("producer is not initialized, call InitProducer first")
		return errors.New("producer is not initialized")
	}

	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(message),
	}
	// 设置延迟级别，这里需要根据 RocketMQ 配置的延迟级别进行调整
	delayLevel := calculateDelayLevel(delay)
	msg.WithDelayTimeLevel(delayLevel)

	res, err := p.SendSync(context.Background(), msg)
	if err != nil {
		mtl.Logger.Error("send delayed message error: %v", err)
		return err
	}
	mtl.Logger.Info("send delayed message success: result=%s", res.String())
	return nil
}

func calculateDelayLevel(delay time.Duration) int {
    // 根据实际 RocketMQ 配置的延迟级别进行映射
    // 例如：1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
    switch {
    case delay <= 1*time.Second:
        return 1
    case delay <= 5*time.Second:
        return 2
    case delay <= 10*time.Second:
        return 3
    case delay <= 30*time.Second:
        return 4
    case delay <= 1*time.Minute:
        return 5
    case delay <= 2*time.Minute:
        return 6
    case delay <= 3*time.Minute:
        return 7
    case delay <= 4*time.Minute:
        return 8
    case delay <= 5*time.Minute:
        return 9
    case delay <= 6*time.Minute:
        return 10
    case delay <= 7*time.Minute:
        return 11
    case delay <= 8*time.Minute:
        return 12
    case delay <= 9*time.Minute:
        return 13
    case delay <= 10*time.Minute:
        return 14
    case delay <= 20*time.Minute:
        return 15
    case delay <= 30*time.Minute:
        return 16
    case delay <= 1*time.Hour:
        return 17
    case delay <= 2*time.Hour:
        return 18
    default:
        return 18  // 如果超过2小时，使用最大延迟级别
    }
}

func ShutdownProducer() error {
	if !producerInitialized {
		return nil
	}
	err := p.Shutdown()
	if err != nil {
		mtl.Logger.Error("shutdown producer error: %v", err)
		return err
	}
	producerInitialized = false
	mtl.Logger.Info("Producer shut down successfully")
	return nil
}
