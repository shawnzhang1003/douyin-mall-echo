package main

import (
	"github.com/MakiJOJO/douyin-mall-echo/pkg/rocketmq"
	"github.com/apache/rocketmq-client-go/v2/consumer"

	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	// product "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/product"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
	"github.com/MakiJOJO/douyin-mall-echo/app/cart/internal/dal"
	product_model "github.com/MakiJOJO/douyin-mall-echo/app/product/model"
)

func rmqInit() {

	rc, err := rocketmq.NewConsumer(&rocketmq.Config{
		NameServerAddr: []string{"127.0.0.1:9876"},
		GroupName:      "OrderServiceConsumerGroup",
	})
	if err != nil {
		panic(err)
	}

	rc.Subscribe("update_product", consumer.MessageSelector{}, messageHandler)

	// 捕获信号，优雅关闭
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Consumer is running. Press Ctrl+C to stop.")
	<-signals // 等待信号

	// 关闭消费者
	fmt.Println("Shutting down consumer...")
	rc.Shutdown()
	fmt.Println("Consumer shut down successfully.")
}

func messageHandler(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	product := product_model.Product{}
	for _, msg := range msgs {
		// fmt.Printf("++++++++Received message: %s\n", string(msg.Body))
		if err := json.Unmarshal([]byte(msg.Body), &product); err != nil {
			fmt.Println("消费者反序列化失败")
		}
		kv_product := dal.Product{
			Id: uint32(product.ID),
			Name: product.Name,
			Description: product.Name,
			Price: product.Price,
		}

		// 将商品信息存入 Redis 缓存，设置过期时间为 1 小时
		productJSON, _ := json.Marshal(kv_product)

		dal.RedisClient.Set(context.Background(), fmt.Sprintf("product:%v", kv_product.Id), string(productJSON), time.Hour).Err()
	}
	return consumer.ConsumeSuccess, nil
}
