package rocketmq

import (
    "context"
    "fmt"
    "github.com/apache/rocketmq-client-go/v2"
    "github.com/apache/rocketmq-client-go/v2/primitive"
    "github.com/apache/rocketmq-client-go/v2/producer"
)

// Producer 定义 RocketMQ 生产者结构体
type Producer struct {
    client rocketmq.Producer
}

// NewProducer 创建一个新的 RocketMQ 生产者实例
func NewProducer(config *Config) (*Producer, error) {
    p, err := rocketmq.NewProducer(
        producer.WithGroupName(config.GroupName),
        producer.WithNsResolver(config.GetNsResolver()),
    )
    if err != nil {
        return nil, fmt.Errorf("创建生产者失败: %v", err)
    }

    if err := p.Start(); err != nil {
        return nil, fmt.Errorf("启动生产者失败: %v", err)
    }

    return &Producer{
        client: p,
    }, nil
}

// SendMessage 发送消息到 RocketMQ
func (p *Producer) SendMessage(ctx context.Context, topic string, body []byte) (*primitive.SendResult, error) {
    msg := &primitive.Message{
        Topic: topic,
        Body:  body,
    }

    result, err := p.client.SendSync(ctx, msg)
    if err != nil {
        return nil, fmt.Errorf("发送消息失败: %v", err)
    }

    return result, nil
}

// Shutdown 关闭生产者
func (p *Producer) Shutdown() error {
    return p.client.Shutdown()
}