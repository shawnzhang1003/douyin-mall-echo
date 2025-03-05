package model

import (
	"github.com/MakiJOJO/douyin-mall-echo/app/order/internal/dal"
	"log"
)

func AutoMigrate() error {
	var allModels = []interface{}{
		&Order{},
		&OrderProduct{},
	}

	log.Printf("Starting auto migration for models: %v\n", allModels)

	if dal.DB == nil {
		log.Printf("dal.DB is null\n")
		return nil
	}

	err := dal.DB.AutoMigrate(allModels...)
	if err != nil {
		log.Printf("Failed to auto migrate tables: %v\n", err)
		return err
	}

	log.Printf("Auto migration completed successfully")
	return nil
}
