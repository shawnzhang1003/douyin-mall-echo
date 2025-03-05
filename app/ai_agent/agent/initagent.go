package agent

import (
	"context"
	"fmt"
	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/agent/models"
	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/agent/tools"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

type Agent struct {
	MyAgent compose.Runnable[[]*schema.Message, []*schema.Message]
}

var GlobalAgent Agent

func InitEino() {
	ctx := context.Background()

	// 初始化 tools ，这里整一个数组，通过不同的初始化方式存放工具
	todoTools := []tool.BaseTool{
		tools.GetProductIdTool(), // 使用 NewTool 方式
		// tools.GetSearchTool(),
		tools.BuyDirectTool(),
		tools.SearchProductsTool(),
	}

	// 创建并配置 ChatModel
	chatModel, err := models.GetArkModel()
	if err != nil {
		fmt.Printf("get chatmodel failed, err=%v", err)
		return
	}

	// 获取工具信息, 用于绑定到 ChatModel
	toolInfos := make([]*schema.ToolInfo, 0, len(todoTools))
	var info *schema.ToolInfo
	for _, todoTool := range todoTools {
		info, err = todoTool.Info(ctx)
		if err != nil {
			fmt.Printf("get ToolInfo failed, err=%v", err)
			return
		}
		toolInfos = append(toolInfos, info)
	}

	// 将 tools 绑定到 ChatModel
	err = chatModel.BindTools(toolInfos)
	if err != nil {
		fmt.Printf("BindTools failed, err=%v", err)
		return
	}

	// 创建 tools 节点
	todoToolsNode, err := compose.NewToolNode(context.Background(), &compose.ToolsNodeConfig{
		Tools: todoTools,
	})
	if err != nil {
		fmt.Printf("NewToolNode failed, err=%v", err)
		return
	}

	// 构建完整的处理链
	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(chatModel, compose.WithNodeName("chat_model")).
		AppendToolsNode(todoToolsNode, compose.WithNodeName("tools"))

	// 编译并运行 chain，生成一个agent
	GlobalAgent.MyAgent, err = chain.Compile(ctx)
	if err != nil {
		fmt.Printf("chain.Compile failed, err=%v", err)
		return
	}
}
