package logic

import (
	"github.com/MakiJOJO/douyin-mall-echo/app/product/model"
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
	product_id, err = model.UpdateProduct(product_id, name, description, picture, price, category_name)
	if err != nil {
		return product_id, err
	}
	return product_id, nil
}

func DeleteProduct(product_id uint32) (err error) {
	err = model.DeleteProduct(product_id)
	if err != nil {
		return err
	}
	return nil
}
