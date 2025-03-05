package model

import "gorm.io/gorm"

//
//  ShopPay
//  @Description: 生成订单号
//

type ShopPay struct {
	gorm.Model
	ByCode          string `gorm:"type:varchar(100)"`
	OrderId         string `gorm:"type:varchar(100); unique;not null"` // 订单ID
	OrderStatus     string `gorm:"type:varchar(100); not null"`        // 订单状态
	OrderTips       string `gorm:"type:varchar(200); not null"`        // 订单备注
	OrderTotalPrice string `gorm:"type:varchar(100); not null"`
	ShopID          string `gorm:"type:varchar(100); not null"`
}
