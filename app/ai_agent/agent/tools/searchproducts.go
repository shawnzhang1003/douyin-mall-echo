package tools

import (
	"context"
	"fmt"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/rpc/client"
)

func SearchProductsTool() tool.InvokableTool {
	info := &schema.ToolInfo{
		Name: "search_products",
		Desc: "search products info by given productinfo",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"productinfo": {
				Desc:     "The name or description of the product",
				Type:     schema.String,
				Required: true,
			},
		}),
	}

	return utils.NewTool(info, SearchProductsFunc)
}

func SearchProductsFunc(_ context.Context, params *SearchProductsReq) (*product.SearchProductsResp, error) {
	fmt.Printf("invoke tool search_products: %+v", params)
	respRPC, err := client.ProductClient.SearchProducts(context.Background(), &product.SearchProductsReq{
		Query: params.ProductInfo,
	})
	if err != nil {
		return nil, err
	}
	return respRPC, nil
}

type SearchProductsReq struct {
	ProductInfo  string `json:"productinfo"`
}