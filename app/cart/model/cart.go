package model

import (
	// "strings"

	"errors"
	"github.com/MakiJOJO/douyin-mall-echo/app/cart/internal/dal"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11);not null;index:idx_user_id" json:"userid"`
	ProductId uint32 `gorm:"type:int(11);not null" json:"productid"`
	Qty       uint32 `gorm:"type:int(11);not null" json:"quantity"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(item *Cart) error {
	var row Cart
	db := dal.DB

	err := db.Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if row.ID > 0 {
		return db.Model(&Cart{}).Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).UpdateColumn("qty", gorm.Expr("qty+?", item.Qty)).Error
	}
	return db.Create(item).Error
}

func GetCartByUserId(userId uint32) ([]*Cart, error) {
	db := dal.DB
	var rows []*Cart
	err := db.Model(&Cart{}).Where(&Cart{UserId: userId}).Find(&rows).Error
	return rows, err
}

func EmptyCart(userid uint32) error {
	db := dal.DB
	if userid == 0 {
		return errors.New("user_id is required")
	}
	return db.Delete(&Cart{}, "user_id = ?", userid).Error
}
