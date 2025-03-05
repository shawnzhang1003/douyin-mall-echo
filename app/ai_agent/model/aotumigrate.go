package model

import "github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/internal/dal"

func AutoMigrate() {
	var allModels = []interface{}{

		// &Cart{},
	}
	dal.DB.AutoMigrate(allModels...)
}
