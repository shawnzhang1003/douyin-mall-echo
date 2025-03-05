package model

import "github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/internal/dal"

func AutoMigrate() {
	var allModels = []interface{}{
	
		&User{},
	
	}
	dal.DB.AutoMigrate(allModels...)
}