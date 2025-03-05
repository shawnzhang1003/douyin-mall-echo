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
	"errors"
	"fmt"

	"github.com/MakiJOJO/douyin-mall-echo/app/product/internal/dal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

func GetById(productId uint32) (product Product, err error) {
	db := dal.DB
	err = db.Model(&Product{}).First(&product, productId).Error
	return product, err
}

func SearchProducts(q string) (products []*Product, err error) {
	db := dal.DB
	err = db.Model(&Product{}).Find(&products, "name like ? or description like ?",
		"%"+q+"%", "%"+q+"%",
	).Error
	return products, err
}

func CreateProduct(productName string, productDescription string, productPicture string, productPrice float32, categoryNames []string) (product *Product, err error) {
	db := dal.DB
	// 检查分类是否存在
	var categories []Category
	for _, categoryName := range categoryNames {
		var category Category
		if err := db.Where("name = ?", categoryName).First(&category).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Printf("Category '%s' not found\n", categoryName)
			}
			fmt.Printf("Error finding category: %v\n", err)
			return nil, err
		}
		categories = append(categories, category)
	}
	// 创建商品
	retProduct := Product{Name: productName, Price: productPrice, Description: productDescription, Picture: productPicture, Categories: categories}
	if err := db.Create(&retProduct).Error; err != nil {
		fmt.Printf("Error creating product: %v\n", err)
		return nil, err
	}

	fmt.Printf("Product '%s' created with categories '%v'\n", productName, categoryNames)
	return &retProduct, nil
}

func UpdateProduct(productId uint32, productName string, productDescription string, productPicture string, productPrice float32, categoryName string) (product_id uint32, err error) {
	db := dal.DB
	result := db.Model(&Product{}).Where("id = ?", productId).Updates(Product{Name: productName, Price: productPrice, Description: productDescription, Picture: productPicture})

	// 检查更新是否成功
	if result.Error != nil {
		return productId, result.Error
	}
	return productId, nil
}

func ListProducts(page int, pageSize int, categoryName string) (products []*Product, err error) {
	if page < 1 || pageSize < 1 {
		return nil, errors.New("invalid page or pageSize")
	}
	db := dal.DB
	var category Category
	if categoryName != "" {
		if err := db.Where("name = ?", categoryName).First(&category).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("category not found")
			}
			return nil, err
		}
	}
	offset := (page - 1) * pageSize

	if err := db.Joins("JOIN product_category ON product_category.product_id = product.id").
		Where("product_category.category_id = ?", category.ID).
		Limit(pageSize).Offset(offset).
		Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func DeleteProduct(productId uint32) (err error) {
	db := dal.DB
	if err := db.Where("id = ?", productId).Delete(&Product{}).Error; err != nil {
		return err
	}
	return nil
}
