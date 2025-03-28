// Code generated by Kitex v0.12.1. DO NOT EDIT.

package ai_agentservice

import (
	"context"
	"errors"
	ai_agent "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/ai_agent"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Hello": kitex.NewMethodInfo(
		helloHandler,
		newHelloArgs,
		newHelloResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	aI_AgentServiceServiceInfo                = NewServiceInfo()
	aI_AgentServiceServiceInfoForClient       = NewServiceInfoForClient()
	aI_AgentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return aI_AgentServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return aI_AgentServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return aI_AgentServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "AI_AgentService"
	handlerType := (*ai_agent.AI_AgentService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "ai_agent",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func helloHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(ai_agent.HelloReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(ai_agent.AI_AgentService).Hello(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *HelloArgs:
		success, err := handler.(ai_agent.AI_AgentService).Hello(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*HelloResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newHelloArgs() interface{} {
	return &HelloArgs{}
}

func newHelloResult() interface{} {
	return &HelloResult{}
}

type HelloArgs struct {
	Req *ai_agent.HelloReq
}

func (p *HelloArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(ai_agent.HelloReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *HelloArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *HelloArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *HelloArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *HelloArgs) Unmarshal(in []byte) error {
	msg := new(ai_agent.HelloReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var HelloArgs_Req_DEFAULT *ai_agent.HelloReq

func (p *HelloArgs) GetReq() *ai_agent.HelloReq {
	if !p.IsSetReq() {
		return HelloArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HelloArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *HelloArgs) GetFirstArgument() interface{} {
	return p.Req
}

type HelloResult struct {
	Success *ai_agent.HelloResp
}

var HelloResult_Success_DEFAULT *ai_agent.HelloResp

func (p *HelloResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(ai_agent.HelloResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *HelloResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *HelloResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *HelloResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *HelloResult) Unmarshal(in []byte) error {
	msg := new(ai_agent.HelloResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HelloResult) GetSuccess() *ai_agent.HelloResp {
	if !p.IsSetSuccess() {
		return HelloResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HelloResult) SetSuccess(x interface{}) {
	p.Success = x.(*ai_agent.HelloResp)
}

func (p *HelloResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HelloResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Hello(ctx context.Context, Req *ai_agent.HelloReq) (r *ai_agent.HelloResp, err error) {
	var _args HelloArgs
	_args.Req = Req
	var _result HelloResult
	if err = p.c.Call(ctx, "Hello", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
