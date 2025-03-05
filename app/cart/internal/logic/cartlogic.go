package logic

import (
	"context"
	"errors"

	"fmt"

	"github.com/MakiJOJO/douyin-mall-echo/app/cart/model"
	"github.com/MakiJOJO/douyin-mall-echo/app/cart/rpc/client"
	product "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/product"
)

func AddItem(item *model.Cart) error {
	// 校验productid是否存在
	_, err := client.ProductClient.GetProduct(context.Background(), &product.GetProductReq{
		Id: item.ProductId,
	})
	if err != nil {
		return errors.New("the product does not exist")
	}

	err = model.AddItem(item)
	if err != nil {
		return err
	}
	return nil
}

func GetCart(userid uint32) (itemList []*model.Cart, err error) {
	if userid == 0 {
		return nil, errors.New("userid can not be zero")
	}
	itemList, err = model.GetCartByUserId(userid)
	if err != nil {
		return nil, err
	}
	// 返回结果, 光return不行，需要c.JSON定义返回内容等函数
	return itemList, nil
}

type retItemInfo struct {
	ProductId   uint32   `protobuf:"varint,1,opt,name=productidid,proto3" json:"productid,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Qty         uint32   `json:"quantity,omitempty"`
	Picture     string   `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	Price       float32  `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Categories  []string `protobuf:"bytes,6,rep,name=categories,proto3" json:"categories,omitempty"`
}

func GetCartReturnInfo(userid uint32) (itemList []*retItemInfo, err error) {
	getCartRet, err := GetCart(userid)

	for _, cartItem := range getCartRet {
		getProductResult, err := client.ProductClient.GetProduct(context.Background(), &product.GetProductReq{Id: cartItem.ProductId})
		if err != nil {
			return nil, err
		}
		if getProductResult.Product == nil {
			return nil, fmt.Errorf("product %v info does not exist", cartItem.ProductId)
		}
		itemList = append(itemList, &retItemInfo{
			ProductId: getProductResult.Product.Id,
			Name: getProductResult.Product.Name,
			Description: getProductResult.Product.Description,
			Qty: cartItem.Qty,
			Picture: getProductResult.Product.Picture,
			Price: getProductResult.Product.Price * float32(cartItem.Qty),
			Categories: getProductResult.Product.Categories,
		})
	}
	return
}

func EmptyCart(userid uint32) error {
	err := model.EmptyCart(userid)
	if err != nil {
		return err
	}
	return nil
}
