package rpc

import (
	"context"
	"testing"

	user "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/user"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

// 修改后的代码
func TestRPC(t *testing.T) {
	registryAddr := []string{"10.21.220.68:2379"} // 改为正确的 etcd 地址

	// 创建 resolver
	r, err := etcd.NewEtcdResolver(registryAddr)
	if err != nil {
		t.Fatalf("create resolver failed: %v", err)
	}

	// 创建客户端
	UserClient, err := userservice.NewClient(
		"user",
		client.WithResolver(r),
		client.WithHostPorts("10.21.220.68:8880"),
	)
	if err != nil {
		t.Fatalf("create client failed: %v", err)
	}

	ctx := context.Background()
	registerReq := &user.RegisterReq{
		Email:           "example23@google.com",
		Password:        "123xyz",
		ConfirmPassword: "123xyz",
	}

	resp, err := UserClient.Register(ctx, registerReq)
	if err != nil {
		t.Fatalf("call register failed: %v", err)
	}
	t.Logf("resp: %v", resp)
}
