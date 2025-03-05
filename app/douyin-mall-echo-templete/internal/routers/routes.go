package routers

import (
	"encoding/json"
	"log"

	"github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/config"
	"github.com/MakiJOJO/douyin-mall-echo/common/middleware"
	"github.com/MakiJOJO/douyin-mall-echo/common/utils"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/health", healthHandler)
	api := e.Group("/api/v1")
	// 下面两个接口不需要jwt鉴权,可以放在这里
	// api.POST("/login", LoginHandler)
	// api.POST("/register", RegisterHandler)
	// 从这之后的路由都会先进行jwt鉴权处理
	api.Use(middleware.JWTAuthMiddleware(config.GlobalConfig.JWT.SecretKey,
		utils.DefaultIsWhiteListFunc(config.GlobalConfig.JWT.Whitelist)))

	{
		// 下面两个接口不需要jwt鉴权, 但是因为前面已经加入了jwt鉴权中间件, 所以需要把path加入config.yaml白名单
		api.POST("/login", LoginHandler)
		api.POST("/register", RegisterHandler)

		api.GET("/hello", HelloWorldHandler)
	}

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		e.Logger.Errorf("Failed to marshal routes: %v", err)
	}
	log.Printf("Routes: %s", string(data))

}
