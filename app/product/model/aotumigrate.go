package model

import "github.com/MakiJOJO/douyin-mall-echo/app/product/internal/dal"

func AutoMigrate() {
	allModels := []interface{}{
		&Product{},
		&Category{},
	}
	dal.DB.AutoMigrate(allModels...)
}
