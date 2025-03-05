package model

import (
	"github.com/MakiJOJO/douyin-mall-echo/app/order/internal/dal"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/sony/sonyflake"
	"gorm.io/gorm"
	"time"
)

// Order 对应 orders 表
type Order struct {
	OrderID     uint64    `gorm:"column:order_id;primaryKey"`
	CartID      int       `gorm:"column:cart_id;not null;unique"`
	UserID      int       `gorm:"column:user_id;not null"`
	UserAddr    string    `gorm:"column:user_addr;not null"`
	OrderStatus string    `gorm:"column:order_status;default:'Pending'"`
	OrderTotal  float64   `gorm:"column:order_total;not null"`
	CreateTime  time.Time `gorm:"column:create_time"`
	CancelTime  time.Time `gorm:"column:cancel_time;default:null"`
	//OrderProducts []OrderProduct `gorm:"foreignKey:OrderID"`
}

// OrderProduct 对应 order_products 表
type OrderProduct struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement"`
	OrderID   uint64  `gorm:"column:order_id;not null"`
	ProductID int     `gorm:"column:product_id;not null"`
	Quantity  int     `gorm:"column:quantity;not null"`
	Price     float64 `gorm:"column:price;not null"`
}

// GetOrderByUser 根据用户 ID 获取订单信息
func GetOrderByUserId(userID int) ([]Order, error) {
	var orders []Order
	result := dal.DB.Where("user_id = ?", userID).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

// GetOrderByUser 根据订单 ID 获取订单信息
func GetOrderByOrderId(orderID uint64) (Order, error) {
	var order Order
	result := dal.DB.Where("order_id = ?", orderID).Find(&order)
	if result.Error != nil {
		return order, result.Error
	}

	return order, nil
}

// GetOrderStatusByOrderId 根据订单 ID 获取订单状态
func GetOrderStatusByOrderId(orderID uint64) (string, error) {
	var order Order
	result := dal.DB.Where("order_id = ?", orderID).First(&order)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return "", nil // 未找到记录，返回空字符串和 nil 错误
		}
		return "", result.Error
	}

	return order.OrderStatus, nil
}

func GetOrderByCartId(cartId int) (Order, error) {
	var order Order
	result := dal.DB.Where("cart_id = ?", cartId).Find(&order)
	if result.Error != nil {
		return order, result.Error
	}

	return order, nil
}

// 创建订单
// Done: 订单幂等性，通过确保一个cartid对应唯一一个订单
// Todo: 熔断限流降级
func CreateOrder(userID int, cartId int, userAddr string, productIDs []int, quantities []int, prices []float64) (uint64, error) {

	order, err := GetOrderByCartId(cartId)
	if err != nil {
		klog.Error(err, "GetOrderByCartId returned an error")
	}
	if order.OrderID != 0 {
		return 0, nil
	}

	startTime := "2024-08-20" // 初始化一个开始的时间，表示从这个时间开始算起
	machineID := 1            // 机器 ID
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		panic(err)
	}

	sonyMachineID := uint16(machineID)
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: func() (uint16, error) {
			return sonyMachineID, nil
		},
	}

	orderId, _ := sonyflake.NewSonyflake(settings).NextID()
	var orderTotal float64
	for i, price := range prices {
		orderTotal += price * float64(quantities[i])
		orderProduct := &OrderProduct{
			OrderID:   orderId,
			ProductID: productIDs[i],
			Quantity:  quantities[i],
			Price:     prices[i],
		}
		result := dal.DB.Create(&orderProduct)
		if result.Error != nil {
			return 0, result.Error
		}
	}

	order = Order{
		OrderID:    orderId,
		CartID:     cartId,
		UserID:     userID,
		UserAddr:   userAddr,
		OrderTotal: orderTotal,
		CreateTime: time.Now(),
	}

	result := dal.DB.Create(&order)
	if result.Error != nil {
		return 0, result.Error
	}

	return orderId, nil
}

// 修改订单信息
// Todo: 引入锁机制，确保同一时间只对于同一订单而言只有一个线程可以修改订单,考虑分布式锁
func UpdateOrderAddr(orderID uint64, userAddr string) error {
	// 构建更新字段的 map
	updates := make(map[string]interface{})
	if userAddr != "" {
		// 确保键名和 Order 结构体中的字段名一致
		updates["UserAddr"] = userAddr
	}

	// 执行更新操作
	result := dal.DB.Model(&Order{}).Where("order_id = ?", orderID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	// 检查是否有记录被更新
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func UpdateOrderTotal(orderID uint64, ordertotal float64) error {
	// 构建更新字段的 map
	updates := make(map[string]interface{})
	if ordertotal >= 0 {
		// 确保键名和 Order 结构体中的字段名一致
		updates["OrderTotal"] = ordertotal
	}

	// 执行更新操作
	result := dal.DB.Model(&Order{}).Where("order_id = ?", orderID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	// 检查是否有记录被更新
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// 定时取消订单
// Done: 引入消息队列做定时取消任务
func CancelOrder(orderID uint64) error {
	result := dal.DB.Model(&Order{}).Where("order_id = ? AND order_status = 'Pending'", orderID).Updates(map[string]interface{}{
		"order_status": "Cancelled",
		"cancel_time":  time.Now(),
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 更新订单状态为成功
func UpdateOrderSuccess(orderID uint64) error {
	// 构建更新字段的 map
	updates := make(map[string]interface{})
	updates["OrderStatus"] = "Success"

	// 执行更新操作
	result := dal.DB.Model(&Order{}).Where("order_id = ?", orderID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	// 检查是否有记录被更新
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
