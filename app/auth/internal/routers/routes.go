package routers

import (
	"github.com/MakiJOJO/douyin-mall-echo/app/auth/config"
	"github.com/MakiJOJO/douyin-mall-echo/common/middleware"
	"github.com/MakiJOJO/douyin-mall-echo/common/utils"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/health", healthHandler)
	api := e.Group("/api/v1")
	api.Any("/auth/refresh", RefreshTokenHandler)

	api.Use(middleware.JWTAuthMiddleware(config.GlobalConfig.JWT.SecretKey,
		utils.DefaultIsWhiteListFunc(config.GlobalConfig.JWT.Whitelist)))
	{
		// api.POST("/login", LoginHandler)
		// api.POST("/register", RegisterHandler)
		api.GET("/hello", HelloWorldHandler)
	}

	// data, err := json.MarshalIndent(e.Routes(), "", "  ")
	// if err != nil {
	// 	e.Logger.Errorf("Failed to marshal routes: %v", err)
	// }
	// log.Printf("Routes: %s", string(data))

}
