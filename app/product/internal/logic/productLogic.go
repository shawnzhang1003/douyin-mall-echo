package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/MakiJOJO/douyin-mall-echo/app/product/model"
	"github.com/MakiJOJO/douyin-mall-echo/pkg/rocketmq"
)

func GetProductsByCategoryID(category_id uint32) (products []*model.Product, err error) {
	products, err = model.GetProductsByCategoryID(category_id)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(product_id uint32) (retproduct model.Product, err error) {
	retproduct, err = model.GetById(product_id)
	if err != nil {
		return model.Product{}, err
	}
	return retproduct, nil
}

func SearchProducts(name string) (products []*model.Product, err error) {
	products, err = model.SearchProducts(name)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetAllCategories() (categories []*model.Category, err error) {
	categories, err = model.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func CreateProduct(name string, description string, picture string, price float32, category_names []string) (retproduct *model.Product, err error) {
	retproduct, err = model.CreateProduct(name, description, picture, price, category_names)
	if err != nil {
		return retproduct, err
	}
	return retproduct, nil
}

func ListProducts(page int, pageSize int, categoryName string) (products []*model.Product, err error) {
	products, err = model.ListProducts(page, pageSize, categoryName)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func UpdateProduct(productID uint32, name string, description string, picture string, price float32, category_name string) (product_id uint32, err error) {
	product, err := model.UpdateProduct(productID, name, description, picture, price, category_name) // 操作数据库
	if err != nil {
		return productID, err
	}

	producer, err := rocketmq.NewProducer(&rocketmq.Config{
		NameServerAddr: []string{"127.0.0.1:9876"},
		GroupName:      "OrderServiceProducerGroup",
	})
	defer producer.Shutdown()
	if err != nil {
		return productID, err
	}
	
	productJSON, err := json.Marshal(product)
	if err != nil {
		return productID, fmt.Errorf("序列化商品信息失败: %v", err)
	}

	_, err = producer.SendMessage(context.Background(), "update_product", productJSON)

	if err != nil {
		return productID, err
	}
	// fmt.Printf("消息发送成功，消息 ID: %s\n", res.MsgID)

	return productID, nil
}

func DeleteProduct(product_id uint32) (err error) {
	err = model.DeleteProduct(product_id)
	if err != nil {
		return err
	}
	return nil
}
