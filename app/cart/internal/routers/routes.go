package routers

import (
	"encoding/json"
	"log"

	"github.com/MakiJOJO/douyin-mall-echo/app/cart/config"
	"github.com/MakiJOJO/douyin-mall-echo/common/middleware"
	"github.com/MakiJOJO/douyin-mall-echo/common/utils"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/health", healthHandler)
	api := e.Group("/api/v1")
	// 以下的接口需要jwt token认证
	api.Use(middleware.JWTAuthMiddleware(config.GlobalConfig.JWT.SecretKey, utils.DefaultIsWhiteListFunc(config.GlobalConfig.JWT.Whitelist)))
	{
		api.POST("/cart", AddItemHandler)
		api.GET("/cart", GetCartHandler)
		api.DELETE("/cart", EmptyCartHandler)
		api.GET("/hello", HelloWorldHandler)
	}

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		e.Logger.Errorf("Failed to marshal routes: %v", err)
	}
	log.Printf("Routes: %s", string(data))

}
