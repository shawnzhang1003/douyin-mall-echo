// Description: This file is used to initialize the rpc client.
package client

import (
	"sync"

	"github.com/MakiJOJO/douyin-mall-echo/app/checkout/config"
	"github.com/MakiJOJO/douyin-mall-echo/common/kitexopt"
	"github.com/cloudwego/kitex/client"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/order/orderservice"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/pay/pay"

	"log"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient pay.Client
	OrderClient   orderservice.Client
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

		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}


func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	if err != nil {
		log.Fatalf("init kitex rpc error: %v", err.Error())
	}
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	if err != nil {
		log.Fatalf("init kitex rpc error: %v", err.Error())
	}
}

func initPaymentClient() {
	PaymentClient, err = pay.NewClient("pay", commonSuite)
	if err != nil {
		log.Fatalf("init kitex rpc error: %v", err.Error())
	}
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	if err != nil {
		log.Fatalf("init kitex rpc error: %v", err.Error())
	}
}

