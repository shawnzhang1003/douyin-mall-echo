// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"github.com/MakiJOJO/douyin-mall-echo/app/product/internal/dal"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `json:"product" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

func GetProductsByCategoryID(categoryID uint32) ([]*Product, error) {
	db := dal.DB
	var products []*Product
	err := db.Joins("JOIN product_category ON product_category.product_id = product.id").
		Where("product_category.category_id = ?", categoryID).
		Find(&products).Error
	return products, err
}

func GetAllCategories() ([]*Category, error) {
	db := dal.DB
	var retCategories []*Category
	err := db.Find(&retCategories).Error
	return retCategories, err
}
