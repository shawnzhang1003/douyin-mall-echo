package rpc

import (
	"context"
	"fmt"
	"github.com/MakiJOJO/douyin-mall-echo/app/order/config"
	"github.com/MakiJOJO/douyin-mall-echo/app/order/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	"os"
	"testing"

	order "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/order"
)

func TestMain(m *testing.M) {
	mtl.InitTracing("order")
	mtl.InitLogger(os.Stdout, "order")
	// 检查文件是否存在
	filePath := "C:\\Users\\NorA\\GolandProjects\\src\\douyin-mall-echo\\app\\order\\config\\config.yaml"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("文件不存在: %s", filePath)
	}
	fmt.Printf("文件存在: %s\n", filePath)

	err := config.Init(filePath)
	if err != nil {
		// 使用 fmt.Sprintf 进行字符串拼接
		panic(fmt.Sprintf("config: %v", err))
	}
	dal.Init()
	if dal.DB == nil {
		fmt.Printf("Test main dal.DB is null!!!!")
	}
	// 执行其他测试函数
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestOrderServiceImpl_CreateOrder(t *testing.T) {
	// 准备测试数据
	req := &order.CreateOrderRequest{
		UserId:     1,
		CartId:     2,
		UserAddr:   "test address",
		ProductIds: []int64{1, 2},
		Quantities: []int64{1, 1},
		Prices:     []float32{10.0, 20.0},
	}

	// 创建 OrderServiceImpl 实例
	s := &OrderServiceImpl{}

	// 调用被测试的方法
	resp, err := s.CreateOrder(context.Background(), req)
	if err != nil {
		t.Errorf("CreateOrder returned an error: %v", err)
		return
	}

	if resp.OrderId <= 0 {
		t.Errorf("Expected a valid order ID, but got %d", resp.OrderId)
	}
}

func TestOrderServiceImpl_OrderPaySuccess(t *testing.T) {
	// 准备测试数据
	req := &order.OrderPaySuccessRequest{
		OrderId: 26755937226194945,
	}

	// 创建 OrderServiceImpl 实例
	s := &OrderServiceImpl{}

	// 调用被测试的方法
	resp, err := s.OrderPaySuccess(context.Background(), req)
	if err != nil {
		t.Errorf("OrderPaySuccess returned an error: %v", err)
		return
	}

	if resp.Message != "order pay success" && resp.Message != "order pay fail" {
		t.Errorf("Expected message 'order pay success' or 'order pay fail', but got '%s'", resp.Message)
	}
}
