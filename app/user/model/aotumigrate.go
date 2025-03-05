package model

import "github.com/MakiJOJO/douyin-mall-echo/app/user/internal/dal"

func AutoMigrate() {
	var allModels = []interface{}{

		&User{},
	}
	dal.DB.AutoMigrate(allModels...)
}
