package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

// Consumer 定义 RocketMQ 消费者结构体
type Consumer struct {
	client rocketmq.PushConsumer
}

// NewConsumer 创建一个新的 RocketMQ 消费者实例
func NewConsumer(config *Config) (*Consumer, error) {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(config.GroupName),
		consumer.WithNsResolver(config.GetNsResolver()),
		consumer.WithConsumeMessageBatchMaxSize(1),
	)
	if err != nil {
		return nil, fmt.Errorf("创建消费者失败: %v", err)
	}

	return &Consumer{
		client: c,
	}, nil
}

// Subscribe 订阅主题并处理消息
func (c *Consumer) Subscribe(topic string, selector consumer.MessageSelector, handler func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)) error {
	if err := c.client.Subscribe(topic, selector, handler); err != nil {
		return fmt.Errorf("订阅主题失败: %v", err)
	}

	if err := c.client.Start(); err != nil {
		return fmt.Errorf("启动消费者失败: %v", err)
	}

	return nil
}

// Shutdown 关闭消费者
func (c *Consumer) Shutdown() error {
	return c.client.Shutdown()
}
