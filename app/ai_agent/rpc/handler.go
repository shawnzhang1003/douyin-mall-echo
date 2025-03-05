package rpc

import (
	"context"
	ai_agent "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/ai_agent"

)

// AI_AgentServiceImpl implements the last service interface defined in the IDL.
type AI_AgentServiceImpl struct{}

// Hello implements the AI_AgentServiceImpl interface.
func (s *AI_AgentServiceImpl) Hello(ctx context.Context, req *ai_agent.HelloReq) (resp *ai_agent.HelloResp, err error) {
	// TODO: Your code here...
	return
}
