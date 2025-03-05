package rpc

import (
	"context"

	"github.com/MakiJOJO/douyin-mall-echo/app/cart/internal/logic"
	"github.com/MakiJOJO/douyin-mall-echo/app/cart/model"
	cart "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/cart"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}

	err = logic.AddItem(cartItem)
	if err != nil {
		return nil, err
	}

	return &cart.AddItemResp{}, nil
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// TODO: Your code here...
	itemList, err := logic.GetCart(req.UserId)
	if err != nil {
		return nil, err
	}

	var items []*cart.CartItem
	for _, v := range itemList {
		items = append(items, &cart.CartItem{
			ProductId: v.ProductId,
			Quantity:  v.Qty,
		})
	}

	return &cart.GetCartResp{Cart: &cart.Cart{
		UserId: req.UserId,
		Items:  items,
	}}, nil
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// TODO: Your code here...
	err = logic.EmptyCart(req.UserId)
	if err != nil {
		return nil, err
	}

	return &cart.EmptyCartResp{}, nil
}
