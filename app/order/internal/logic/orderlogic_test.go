package logic

import (
	"fmt"
	"github.com/MakiJOJO/douyin-mall-echo/app/order/config"
	"github.com/MakiJOJO/douyin-mall-echo/app/order/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/app/order/model"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	"os"
	"testing"
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

func TestAutoMigrate(t *testing.T) {
	if dal.DB == nil {
		fmt.Printf("dal.DB is null!!!!")
	}
	model.AutoMigrate()
}

func TestCreateOrder(t *testing.T) {

	userId := 1155666
	cartId := 5698887
	userAddr := "secret garden 5566"
	productIDs := []int{1, 2}
	quantities := []int{2, 3}
	prices := []float64{10.0, 20.0}

	_, err := model.CreateOrder(userId, cartId, userAddr, productIDs, quantities, prices)
	if err != nil {
		panic(err)
	}
}

func TestGetOrder(t *testing.T) {
	orderid := uint64(26142364875620353)
	order, err := model.GetOrderByOrderId(orderid)
	if err != nil {
		t.Errorf("GetOrderByOrderId returned an error: %v", err)
		return
	}

	// 直接输出 order 结构体
	fmt.Printf("Order: %+v\n", order)
}

func TestUpdateOrderAddr(t *testing.T) {
	orderid := uint64(26750396064071681)
	useraddr := "lake behind"
	err := model.UpdateOrderAddr(orderid, useraddr)
	if err != nil {
		t.Errorf("GetOrderByOrderId returned an error: %v", err)
		return
	}

}

func TestUpdateOrderTotal(t *testing.T) {
	orderid := uint64(26750396064071681)
	ordertotal := 6677.44
	err := model.UpdateOrderTotal(orderid, ordertotal)
	if err != nil {
		t.Errorf("GetOrderByOrderId returned an error: %v", err)
		return
	}

}

func TestGetOrderStatusByOrderId(t *testing.T) {
	orderid := uint64(26750396064071681)
	Orderstatus, err := model.GetOrderStatusByOrderId(orderid)
	if err != nil {
		t.Errorf("GetOrderByOrderId returned an error: %v", err)
	}
	fmt.Printf("Orderstatus: %+v\n", Orderstatus)
}

func TestGetOrderTotalByCartId(t *testing.T) {
	cartid := 5698887
	order, err := model.GetOrderByCartId(cartid)
	if err != nil {
		t.Errorf("GetOrderByOrderId returned an error: %v", err)
		return
	}

	// 直接输出 order 结构体
	fmt.Printf("Order: %+v\n", order)
}

func TestUpdateOrderSuccess(t *testing.T) {
	orderid := uint64(26750396064071681)
	err := model.UpdateOrderSuccess(orderid)
	if err != nil {
		t.Errorf("GetOrderByOrderId returned an error: %v", err)
		return
	}

}
