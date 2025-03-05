package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/MakiJOJO/douyin-mall-echo/app/pay/config"
	"github.com/MakiJOJO/douyin-mall-echo/app/pay/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/app/pay/internal/routers"
	"github.com/MakiJOJO/douyin-mall-echo/app/pay/internal/server"
	"github.com/MakiJOJO/douyin-mall-echo/app/pay/rpc"

	rpcclient "github.com/MakiJOJO/douyin-mall-echo/app/pay/rpc/client"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"

	// "github.com/MakiJOJO/douyin-mall-echo/common/kitexopt"
	// "github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	// "github.com/MakiJOJO/douyin-mall-echo/common/utils"
	pay "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/pay/pay"
	"golang.org/x/net/http2"
)

func main() {
	// 初始化配置
	config.Init("../config/config.yaml")
	serviceName := config.GlobalConfig.Kitex.Service
	serveceAddr := config.GlobalConfig.Kitex.Address
	println("kitex service name: ", serveceAddr)
	rpcclient.InitClient()
	// trace
	mtl.InitTracing(serviceName)
	// metric
	mtl.InitMetric(serviceName, config.GlobalConfig.Kitex.MetricsPort, config.GlobalConfig.Registry.RegistryAddress[0])

	// 初始化mysql
	dal.Init()
	// 自动迁移
	// model.AutoMigrate()
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, addr := range addrs {
		// 这个网络地址是IP地址：IPv4，IPv6
		ipnet, ok := addr.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			fmt.Println("Current IP:", ipnet.IP.String())
			break
		}
	}
	server := server.NewEchoServer()
	// 注册路由
	routers.RegisterRoutes(server)

	// Create a done channel to signal when the shutdown is complete
	done := make(chan struct{}, 1)
	// defer func ()  {
	// 	// Wait for the graceful shutdown to complete
	// 	<-done
	// 	log.Println("Graceful shutdown complete.")
	// }()

	// Run graceful shutdown in a separate goroutine
	// 创建rpc server服务实例并自动连接etcd进行服务注册
	opts := kitexInit()
	// 创建user服务的rpc server, 需要传入实现了user.UserService接口的实例
	// 不同的微服务要修改成对应的rpc服务实例
	kitexSvr := pay.NewServer(new(rpc.PayImpl), opts...)
	go func() {
		// Run the rpc server
		err := kitexSvr.Run()
		if err != nil {
			log.Print(err)
		}
	}()

	go gracefulShutdown(server, kitexSvr, done)

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	if err := server.StartH2CServer(config.GlobalConfig.HOST, s); err != http.ErrServerClosed {
		server.Logger.Fatal("http server error: %s", err)
	}

	// Wait for the graceful shutdown to complete
	<-done
	server.Logger.Info("Graceful shutdown complete.")
}
