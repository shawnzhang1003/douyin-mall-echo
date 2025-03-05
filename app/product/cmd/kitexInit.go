package main

import (
	"net"
	"strings"

	"github.com/MakiJOJO/douyin-mall-echo/app/product/config"
	"github.com/MakiJOJO/douyin-mall-echo/common/kitexopt"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	"github.com/MakiJOJO/douyin-mall-echo/common/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
)

func kitexInit() (opts []server.Option) {
	// address
	address := config.GlobalConfig.Kitex.Address
	if strings.HasPrefix(address, ":") {
		localIp := utils.MustGetLocalIPv4()
		address = localIp + address
	}
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	_ = provider.NewOpenTelemetryProvider(
		provider.WithSdkTracerProvider(mtl.TracerProvider),
		provider.WithEnableMetrics(false),
	)
	serviceName := config.GlobalConfig.Kitex.Service

	opts = append(opts, server.WithSuite(kitexopt.CommonServerSuite{
		CurrentServiceName: serviceName,
		RegistryAddr:       config.GlobalConfig.Registry.RegistryAddress,
		RegistryUsername:   config.GlobalConfig.Registry.Username,
		RegistryPassword:   config.GlobalConfig.Registry.Password,
	}))
	return
}
