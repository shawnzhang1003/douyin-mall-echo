package rocketmq

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/MakiJOJO/douyin-mall-echo/app/order/config"
	"github.com/MakiJOJO/douyin-mall-echo/app/order/model"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func CreateTopic() error {
	// 创建主题
	// 先连接远程的服务器，得到一个具柄testAdmin，然后利用该具柄创建CreateTopic()创建topic
	Admin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(config.GlobalConfig.RocketMq.NsAddrs)))
	// 检查是否连接成功
	if err != nil {
		mtl.Logger.Error("connection mq error", "mqerr", err.Error())
		return err
	}

	// 创建主题
	err = Admin.CreateTopic(context.Background(),
		admin.WithTopicCreate(config.GlobalConfig.RocketMq.Topic),
		admin.WithBrokerAddrCreate(config.GlobalConfig.RocketMq.BrAddrs[0]),
	)

	// 检查是否创建topic失败
	if err != nil {
		mtl.Logger.Error("createTopic error!", "mqerr", err.Error())
		return err
	}

	// 等待一段时间，让 NameServer 更新元数据
	time.Sleep(5 * time.Second)

	return nil
}

// orderTimeout 处理订单超时消息
func orderTimeout(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for _, message := range msgs {
		orderidStr := string(message.Body)
		orderid, err := strconv.ParseUint(orderidStr, 10, 64)
		if err != nil {
			mtl.Logger.Error("Failed to parse order ID from message", "err", err)
			continue
		}

		// 设置上下文超时时间
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		orderStatus, err := model.GetOrderStatusByOrderId(orderid)
		if err != nil {
			mtl.Logger.Error("Failed to get order status for order ID.", "orderid", orderid, "err", err.Error())
			continue
		}

		if orderStatus != "success" {
			err = model.CancelOrder(orderid)
			if err != nil {
				mtl.Logger.Error("Failed to cancel order with order ID .", "orderid", orderid, "err", err.Error())
			}
		}
	}
	return consumer.ConsumeSuccess, nil
}

// InitOrderTimeoutConsumer 初始化订单超时消费者
func InitOrderTimeoutConsumer(ctx context.Context) error {
	config := ConsumerConfig{
		config.GlobalConfig.RocketMq,
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			return orderTimeout(ctx, msgs...)
		},
	}

	err := InitConsumer(ctx, config)
	if err != nil {
		mtl.Logger.Error("failed to initialize order timeout consumer", "consumerErr", err)
		return errors.New("failed to initialize order timeout consumer")
	}
	return nil
}

func Init() {
	// 创建上下文 初始化rocketmq
	ctx := context.Background()

	// 创建 Topic
	err := CreateTopic()
	if err != nil {
		mtl.Logger.Error("Failed to create topic", "mqerr", err)
	}

	// 初始化消费者
	err = InitOrderTimeoutConsumer(ctx)
	if err != nil {
		mtl.Logger.Error("Failed to initialize order timeout consumer", "consumerErr", err)
	}

	// 初始化生产者
	err = InitProducer()
	if err != nil {
		mtl.Logger.Error("Failed to initialize producer", "producerErr", err)

	}
	//defer ShutdownProducer()
}
