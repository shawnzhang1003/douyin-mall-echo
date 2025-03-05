package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/config"
	"github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/internal/routers"
	"github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/internal/server"
	"github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/rpc"
	rpcclient "github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/rpc/client"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"

	// "github.com/MakiJOJO/douyin-mall-echo/common/kitexopt"
	// "github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	// "github.com/MakiJOJO/douyin-mall-echo/common/utils"
	user "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/user/userservice"
	"golang.org/x/net/http2"
)

func main() {
	// 初始化配置
	config.Init("../config/config.yaml")
	var serviceName = config.GlobalConfig.Kitex.Service
	rpcclient.InitClient()
	// trace
	mtl.InitTracing(serviceName)
	// metric
	mtl.InitMetric(serviceName, config.GlobalConfig.Kitex.MetricsPort, config.GlobalConfig.Registry.RegistryAddress[0])

	// 初始化mysql以及自动迁移, 两个一起使用
	// dal.Init()
	// model.AutoMigrate()

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
	kitexSvr := user.NewServer(new(rpc.UserServiceImpl), opts...)
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
