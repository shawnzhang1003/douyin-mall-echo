package rpc

import (
	"context"
	"fmt"
	"github.com/MakiJOJO/douyin-mall-echo/app/checkout/internal/logic"
	checkout "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/checkout"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct{}

// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	address := req.GetAddress()
	// 调用逻辑层
	retInfo, err := logic.Checkout(req.UserId,
		fmt.Sprintf("%s, %s, %s, %s, %s",
			address.GetStreetAddress(),
			address.GetCity(),
			address.GetState(),
			address.GetCountry(),
			address.GetZipCode()))
	resp = &checkout.CheckoutResp{
		OrderId: retInfo.OrderId,
		PayUrl:  retInfo.PayUrl,
	}
	fmt.Println("retUrl:", resp.PayUrl)
	return resp, nil
}
