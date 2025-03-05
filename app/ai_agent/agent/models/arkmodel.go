package models

import (
	"context"
	"fmt"
	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/config"
	"github.com/cloudwego/eino-ext/components/model/ark"
)

func GetArkModel() (*ark.ChatModel, error){
	// 创建并配置 ChatModel
	chatModel, err := ark.NewChatModel(context.Background(), &ark.ChatModelConfig{
		APIKey: config.GlobalConfig.ApiKey, // api key
		Model:  config.GlobalConfig.Model,  // 模型名称
	})
	if err != nil {
		fmt.Printf("NewChatModel failed, err=%v", err)
		return nil, err
	}
	return chatModel, err
}
