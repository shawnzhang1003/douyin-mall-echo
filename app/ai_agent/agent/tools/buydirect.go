package tools

import (
	"context"
	"fmt"

	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/rpc/client"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/order"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/pay"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
	uuid "github.com/google/uuid"
)

func BuyDirectTool() tool.InvokableTool {
	info := &schema.ToolInfo{
		Name: "buy_products",
		Desc: "help user order products by product info",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"productinfo": {
				Desc:     "The product info",
				Type:     schema.String,
				Required: true,
			},
			"userid": {
				Desc:     "The id of the user",
				Type:     schema.Integer,
				Required: true,
			},
			"quantity": {
				Desc:     "The quantity of the product",
				Type:     schema.Integer,
				Required: true,
			},
		}),
	}

	return utils.NewTool(info, BuyDirectFunc)
}

func BuyDirectFunc(_ context.Context, params *BuyDirectReq) (*BuyDirectResp, error) {
	fmt.Printf("invoke tool buy_products: %+v", params)
	respSearchProduct, err := client.ProductClient.SearchProducts(context.Background(), &product.SearchProductsReq{
		Query: params.ProductInfo,
	})
	if err != nil {
		return nil, err
	}

	var productIds []int64
	var quantities []int64
	var prices []float32
	var totalPrice float64

	for _, singleProduct := range respSearchProduct.Results {
		productIds = append(productIds, int64(singleProduct.Id))
		quantities = append(quantities, int64(params.Quantity))
		cost := float32(singleProduct.Price) * float32(params.Quantity)
		prices = append(prices, singleProduct.Price)
		totalPrice += float64(cost)
	}

	respCreateOrder, err := client.OrderClient.CreateOrder(context.Background(), &order.CreateOrderRequest{
		UserId:   int64(params.UserID),
		CartId:   int64(uuid.New().ID()),
		UserAddr: "test addr",
		ProductIds: productIds,
		Quantities: quantities,
		Prices: prices,
	})
	if err != nil {
		return nil, err
	}

	respAliPay, err := client.PaymentClient.AliPay(context.Background(), &pay.PayRequest{
		OrderId:     fmt.Sprintf("%d", respCreateOrder.OrderId),
		TotoalPrice: totalPrice,
	})

	if err != nil {
		return nil, err
	}

	return &BuyDirectResp{
		OrderID:    respCreateOrder.OrderId,
		PayUrl:     respAliPay.PayUrl,
		TotalPrice: totalPrice,
	}, nil
}

type BuyDirectReq struct {
	ProductInfo  string `json:"productinfo"`
	UserID   uint32 `json:"userid"`
	Quantity uint32 `json:"quantity"`
}

type BuyDirectResp struct {
	OrderID    uint64  `json:"orderid"`
	PayUrl     string  `json:"payurl"`
	TotalPrice float64 `json:"totalprice"`
}
