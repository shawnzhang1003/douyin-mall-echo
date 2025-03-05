package model

import "github.com/MakiJOJO/douyin-mall-echo/app/pay/internal/dal"

func AutoMigrate() {
	allModels := []interface{}{
		&ShopPay{},
	}
	dal.DB.AutoMigrate(allModels...)
}
