package model

import "github.com/MakiJOJO/douyin-mall-echo/app/checkout/internal/dal"

func AutoMigrate() {
	var allModels = []interface{}{

		// &User{},
	}
	dal.DB.AutoMigrate(allModels...)
}
