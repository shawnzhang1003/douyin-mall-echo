package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/MakiJOJO/douyin-mall-echo/app/checkout/rpc/client"
	cart "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/cart"
	order "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/order"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/pay"
	product "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/product"
	uuid "github.com/google/uuid"
)

type CheckoutReq struct {
	UserId    uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Firstname string   `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string   `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email     string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Address   *Address `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

type Address struct {
	StreetAddress string `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country       string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode       string `protobuf:"bytes,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
}

type CheckoutRetInfo struct {
	OrderId uint64
	PayUrl  string
}

func Checkout(userid uint32, userAddr string) (retInfo *CheckoutRetInfo, err error) {
	// 获取购物车物品
	getCartResult, err := client.CartClient.GetCart(context.Background(), &cart.GetCartReq{UserId: userid})
	if err != nil {
		return nil, err
	}

	if getCartResult == nil || getCartResult.Cart == nil {
		return nil, errors.New("cart is empty")
	}
	var total float32

	// 从产品服务那里获取价格，算某种商品的总价，总价交给订单服务去算
	var productIds []int64
	var productQuantities []int64
	var productPrices []float32

	for _, cartItem := range getCartResult.Cart.Items {
		getProductResult, err := client.ProductClient.GetProduct(context.Background(), &product.GetProductReq{Id: cartItem.ProductId})
		if err != nil {
			return nil, err
		}
		if getProductResult.Product == nil {
			return nil, fmt.Errorf("product %v info does not exist", cartItem.ProductId)
		}

		productIds = append(productIds, int64(cartItem.ProductId))
		productQuantities = append(productQuantities, int64(cartItem.Quantity))
		cost := getProductResult.Product.Price * float32(cartItem.Quantity)
		productPrices = append(productPrices, getProductResult.Product.Price)
		total += cost
	}

	if total < 0 {
		total = 0
	}

	// for _, id := range productIds {
	// 	fmt.Println("product id: ", id)
	// }

	// 调用订单，创建订单
	retOrderId, err := client.OrderClient.CreateOrder(context.Background(), &order.CreateOrderRequest{
		UserId:     int64(userid),
		CartId:     int64(uuid.New().ID()), //购物车没有分表，暂时返回uuid以hack
		UserAddr:   userAddr,
		ProductIds: productIds,
		Quantities: productQuantities,
		Prices:     productPrices,
	})
	if err != nil {
		return nil, err
	}

	// 清空购物车
	_, err = client.CartClient.EmptyCart(context.Background(), &cart.EmptyCartReq{
		UserId: userid,
	})
	if err != nil {
		return nil, err
	}

	//调用支付
	payResult, err := client.PaymentClient.AliPay(context.Background(), &pay.PayRequest{
		OrderId:     fmt.Sprintf("%d", retOrderId.OrderId),
		TotoalPrice: float64(total),
	})
	if err != nil {
		return nil, err
	}

	// //支付回调，更新order状态
	// client.OrderClient.OrderPaySuccess(context.Background(), &order.OrderPaySuccessRequest{
	// 	OrderId: retOrderId.OrderId,
	// })

	return &CheckoutRetInfo{
		OrderId: retOrderId.OrderId,
		PayUrl:  payResult.PayUrl,
	}, nil
}
