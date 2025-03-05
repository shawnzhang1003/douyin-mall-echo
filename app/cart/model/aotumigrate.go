package model

import "github.com/MakiJOJO/douyin-mall-echo/app/cart/internal/dal"

func AutoMigrate() {
	var allModels = []interface{}{

		&Cart{},
	}
	dal.DB.AutoMigrate(allModels...)
}
