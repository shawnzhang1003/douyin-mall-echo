package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	// rpcclient "github.com/MakiJOJO/douyin-mall-echo/app/auth/rpc/client"
	"github.com/MakiJOJO/douyin-mall-echo/app/user/config"
	"github.com/MakiJOJO/douyin-mall-echo/app/user/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/app/user/internal/routers"
	"github.com/MakiJOJO/douyin-mall-echo/app/user/internal/server"
	"github.com/MakiJOJO/douyin-mall-echo/app/user/model"
	"github.com/MakiJOJO/douyin-mall-echo/app/user/rpc"
	rpcclient "github.com/MakiJOJO/douyin-mall-echo/app/user/rpc/client"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	user "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/user/userservice"

	"golang.org/x/net/http2"
)

func main() {

	// 初始化配置
	config.Init("./app/user/config/config.yaml")
	var serviceName = config.GlobalConfig.Kitex.Service
	rpcclient.InitClient()
	mtl.InitTracing(serviceName)
	mtl.InitMetric(serviceName, config.GlobalConfig.Kitex.MetricsPort, config.GlobalConfig.Registry.RegistryAddress[0])

	// 初始化mysql
	dal.Init()
	// 自动迁移
	model.AutoMigrate()

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
	kitexSvr := user.NewServer(new(rpc.UserServiceImpl), opts...)
	go func() {
		// Run the rpc server
		fmt.Printf("run\n")
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
