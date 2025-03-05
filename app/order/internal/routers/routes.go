package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MakiJOJO/douyin-mall-echo/app/order/config"
	"github.com/MakiJOJO/douyin-mall-echo/app/order/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/common/middleware"
	"github.com/MakiJOJO/douyin-mall-echo/common/utils"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", healthHandler)
	api := e.Group("/api/order")
	// 以下的接口需要jwt token认证
	api.Use(middleware.JWTAuthMiddleware(config.GlobalConfig.JWT.SecretKey, utils.DefaultIsWhiteListFunc(config.GlobalConfig.JWT.Whitelist)))
	{
		api.GET("/getOrderbyUserId", GetOrderByUserIDHandler)
		api.GET("/getOrderbyOrderId", GetOrderByOrderIdHandler)
		api.POST("/create", CreateOrderHandler)
		api.POST("/updateOrderAddr", UpdateOrderAddrHandler)
		api.POST("/cancel", CancelOrderHandler)
	}

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		e.Logger.Errorf("Failed to marshal routes: %v", err)
	}
	log.Printf("Routes: %s", string(data))

}

func healthHandler(c echo.Context) error {
	if dal.DbInstance == nil {
		return c.JSON(http.StatusInternalServerError, "database is not connected")
	}
	return c.JSON(http.StatusOK, dal.DbInstance.Health())
}
