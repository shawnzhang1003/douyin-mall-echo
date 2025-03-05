package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/config"
	// "github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/model"
	// "github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/internal/routers"
	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/internal/server"
	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/rpc"
	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/rpc/client"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	ai_agent "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/ai_agent/ai_agentservice"
	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/agent"


	"golang.org/x/net/http2"
)

func main() {

	// 初始化配置
	config.Init()
	var serviceName = config.GlobalConfig.Kitex.Service
	mtl.InitTracing(serviceName)
	mtl.InitMetric(serviceName, config.GlobalConfig.Kitex.MetricsPort, config.GlobalConfig.Registry.RegistryAddress[0])

	agent.InitEino()

	// 初始化mysql
	// dal.Init()
	// 自动迁移
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

	kitexSvr := ai_agent.NewServer(new(rpc.AI_AgentServiceImpl), opts...)

	client.InitClient()

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
		log.Fatalf("http server error: %s", err)
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")

}
