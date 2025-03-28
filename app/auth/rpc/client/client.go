// Description: This file is used to initialize the rpc client.
package client

import (
	"sync"

	"github.com/MakiJOJO/douyin-mall-echo/app/auth/config"
	"github.com/MakiJOJO/douyin-mall-echo/common/kitexopt"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var (
	// CartClient    cartservice.Client
	// ProductClient productcatalogservice.Client
	// PaymentClient paymentservice.Client
	UserClient   userservice.Client
	once         sync.Once
	err          error
	registryAddr []string
	serviceName  string
	commonSuite  client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = config.GlobalConfig.Registry.RegistryAddress
		serviceName = config.GlobalConfig.Kitex.Service
		commonSuite = client.WithSuite(kitexopt.CommonGrpcClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr,
		})
		// 只初始化本服务需要远程调用的rpc服务, 不同的rpc服务需要创建不同的client
		initUserClient()
		// initCartClient()
		// initProductClient()
		// initPaymentClient()
		// initOrderClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	if err != nil {
		mtl.Logger.Error("init user kitex rpc error.", "err: ", err.Error())
	}
}
