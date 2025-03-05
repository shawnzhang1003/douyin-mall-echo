package model

import "github.com/MakiJOJO/douyin-mall-echo/app/auth/internal/dal"

func AutoMigrate() {
	var allModels = []interface{}{

		&User{},
	}
	dal.DB.AutoMigrate(allModels...)
}
