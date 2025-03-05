package logic

import (
	"fmt"
	"time"

	"github.com/MakiJOJO/douyin-mall-echo/app/order/model"
	"github.com/MakiJOJO/douyin-mall-echo/app/order/rocketmq"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
)

// Order 定义订单结构体
type Order struct {
	ID     uint
	UserID uint
	// 其他订单字段
}

// GetOrderByUserID 根据用户 ID 获取订单信息
func GetOrderByUserID(userID int) ([]model.Order, error) {
	orders, err := model.GetOrderByUserId(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders by user ID %d: %w", userID, err)
	}
	return orders, nil
}

// GetOrderByOrderId 根据订单 ID 获取订单信息
func GetOrderByOrderId(orderID uint64) (model.Order, error) {
	order, err := model.GetOrderByOrderId(orderID)
	if err != nil {
		return model.Order{}, fmt.Errorf("failed to get order by order ID %d: %w", orderID, err)
	}
	return order, nil
}

// CreateOrder 创建订单，并发送延迟消息
func CreateOrder(userid int, cartid int, useraddr string, productids []int, quantities []int, prices []float64) (uint64, error) {
	orderID, err := model.CreateOrder(userid, cartid, useraddr, productids, quantities, prices)
	if err != nil {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	topic := "order_timeout_check"
	message := fmt.Sprintf("%d", orderID)
	err = rocketmq.SendDelayedMessage(topic, message, 30*time.Minute)
	if err != nil {
		mtl.Logger.Error("Failed to send delayed message for order ID ", "orderid", orderID, "err", err.Error())
	}

	return orderID, nil
}

// UpdateOrderAddr 更新订单地址
func UpdateOrderAddr(orderid uint64, useraddr string) error {
	status, err := model.GetOrderStatusByOrderId(orderid)
	if err != nil {
		return fmt.Errorf("failed to update order address get order ID fail %d: %w", orderid, err)
	}
	if status == "Success" {
		return fmt.Errorf("fail to update order because order is paied order ID %d", orderid)
	}
	err = model.UpdateOrderAddr(orderid, useraddr)
	if err != nil {
		return fmt.Errorf("failed to update order address for order ID %d: %w", orderid, err)
	}
	return nil
}

// UpdateOrderAddr 更新订单地址
func UpdateOrderTotal(orderid uint64, ordertotal float64) error {
	status, err := model.GetOrderStatusByOrderId(orderid)
	if err != nil {
		return fmt.Errorf("failed to update order address get order ID fail %d: %w", orderid, err)
	}
	if status == "Success" {
		return fmt.Errorf("fail to update order because order is paied order ID %d", orderid)
	}
	err = model.UpdateOrderTotal(orderid, ordertotal)
	if err != nil {
		return fmt.Errorf("failed to update order address for order ID %d: %w", orderid, err)
	}
	return nil
}

// CancelOrder 取消订单
func CancelOrder(orderid uint64) error {
	err := model.CancelOrder(orderid)
	if err != nil {
		return fmt.Errorf("failed to cancel order with order ID %d: %w", orderid, err)
	}
	return nil
}

// UpdateOrderSuccess 更新订单状态为成功
func UpdateOrderSuccess(orderid uint64) error {
	err := model.UpdateOrderSuccess(orderid)
	if err != nil {
		return fmt.Errorf("failed to update order status to success for order ID %d: %w", orderid, err)
	}
	return nil
}
