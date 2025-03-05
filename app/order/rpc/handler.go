package rpc

import (
	"context"

	"github.com/MakiJOJO/douyin-mall-echo/app/order/internal/logic"
	order "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/order"
)

type OrderServiceImpl struct{}

func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (resp *order.CreateOrderResponse, err error) {
	userid := req.GetUserId()
	cartid := req.GetCartId()
	useraddr := req.GetUserAddr()
	productids64 := req.GetProductIds()
	quantities64 := req.GetQuantities()
	pricesf32 := req.GetPrices()

	// 将 []int64 转换为 []int
	productids := make([]int, len(productids64))
	for i, v := range productids64 {
		productids[i] = int(v)
	}

	quantities := make([]int, len(quantities64))
	for i, v := range quantities64 {
		quantities[i] = int(v)
	}

	prices := make([]float64, len(pricesf32))
	for i, v := range pricesf32 {
		prices[i] = float64(v)
	}

	orderid, err := logic.CreateOrder(int(userid), int(cartid), useraddr, productids, quantities, prices)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{OrderId: orderid}, nil
}

// Login implements the UserServiceImpl interface.
func (s *OrderServiceImpl) OrderPaySuccess(ctx context.Context, req *order.OrderPaySuccessRequest) (resp *order.OrderPaySuccessResponse, err error) {
	orderid := req.GetOrderId()
	err = logic.UpdateOrderSuccess(orderid)
	if err != nil {
		return &order.OrderPaySuccessResponse{Message: "order pay fail"}, err
	}
	return &order.OrderPaySuccessResponse{Message: "order pay success"}, err
}
