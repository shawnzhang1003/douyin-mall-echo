package logic

import (
	"context"

	"fmt"
	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/agent"
	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/agent/prompts"


	// "github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/rpc/client"
	"github.com/cloudwego/eino/schema"
)

func AgentInvoke(ctx context.Context, userid uint32, query []byte) (any, error) {
	// 调用初始化好的agent
	resp, err := agent.GlobalAgent.MyAgent.Invoke(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content:  prompts.GetPrompt() + string(append([]byte(fmt.Sprintf("我是用户 %d, ", userid)), query...)),
		},
	})
	if err != nil {
		fmt.Printf("agent.Invoke failed, err=%v", err)
		return nil, err
	}

	// 输出结果
	for idx, msg := range resp {
		fmt.Printf("\n")
		fmt.Printf("message %d: %s: %s", idx, msg.Role, msg.Content)
	}
	return resp, nil
}
