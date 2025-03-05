package rpc

import (
	"context"
	"strconv"

	"github.com/MakiJOJO/douyin-mall-echo/app/pay/internal/logic"
	pay "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/pay"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// PayImpl implements the last service interface defined in the IDL.
type PayImpl struct{}

// AliPay implements the PayImpl interface.
func (s *PayImpl) AliPay(ctx context.Context, req *pay.PayRequest) (resp *pay.PayResponse, err error) {
	// TODO: Your code here...
	orderId := req.GetOrderId()
	if orderId == "" {
		return nil, kerrors.NewGRPCBizStatusError(2005001, "order id is required")
	}
	totoalPrice := req.GetTotoalPrice()
	if totoalPrice <= 0 {
		return nil, kerrors.NewGRPCBizStatusError(2005002, "total price is invalid")
	}
	priceStr := strconv.FormatFloat(totoalPrice, 'f', 2, 64)
	retpayUrl, err := logic.AliPay(orderId, priceStr)
	if retpayUrl == "" {
		return nil, kerrors.NewGRPCBizStatusError(2005003, "pay url is empty")
	}
	resp = &pay.PayResponse{
		PayUrl: retpayUrl,
	}
	return resp, nil
}
