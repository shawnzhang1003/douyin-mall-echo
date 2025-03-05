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

func GetProductIdTool() tool.InvokableTool {
	info := &schema.ToolInfo{
		Name: "get_product",
		Desc: "Get a product info by id",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"id": {
				Desc:     "The id of the product",
				Type:     schema.Integer,
				Required: true,
			},
		}),
	}

	return utils.NewTool(info, GetProductIdFunc)
}

func GetProductIdFunc(_ context.Context, params *product.GetProductReq) (*product.GetProductResp, error) {
	fmt.Printf("invoke tool get_product: %+v", params)
	respRPC, err := client.ProductClient.GetProduct(context.Background(), params)
	if err != nil {
		return nil, err
	}
	return respRPC, nil
}






