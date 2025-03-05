package routers

import (
	"encoding/json"
	"log"

	"github.com/MakiJOJO/douyin-mall-echo/app/pay/config"
	"github.com/MakiJOJO/douyin-mall-echo/app/pay/utils"
	"github.com/MakiJOJO/douyin-mall-echo/common/middleware"
	commonutils "github.com/MakiJOJO/douyin-mall-echo/common/utils"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", healthHandler)
	api := e.Group("/api/v1")
	// 以下的接口需要jwt token认证
	api.Use(middleware.JWTAuthMiddleware(config.GlobalConfig.JWT.SecretKey, commonutils.DefaultIsWhiteListFunc(config.GlobalConfig.JWT.Whitelist)))

	{
		api.GET("/pay", AliPayHandler)
		api.POST("/pay/notify", utils.Notify)
		api.GET("/pay/callback", utils.Callback)
	}

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		e.Logger.Errorf("Failed to marshal routes: %v", err)
	}
	log.Printf("Routes: %s", string(data))
}
