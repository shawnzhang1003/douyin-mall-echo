package tools

import (
	"fmt"
	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino/components/tool"
	"context"
)

func GetSearchTool() tool.InvokableTool {
	// 创建 Google Search 工具
	searchTool, err := duckduckgo.NewTool(context.Background(), &duckduckgo.Config{})
	if err != nil {
		fmt.Printf("NewGoogleSearchTool failed, err=%v", err)
		return nil
	}
	return searchTool
}