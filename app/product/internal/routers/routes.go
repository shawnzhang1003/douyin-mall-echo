package routers

import (
	"encoding/json"
	"log"

	"github.com/MakiJOJO/douyin-mall-echo/app/product/config"
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
		_product := api.Group("/product")
		_product.GET("/getProductsByCategoryID", GetProductsByCategoryIDHandler)
		_product.GET("/getProductByID", GetProductByIDHandler)
		_product.GET("/searchProduct", SearchProductsHandler)
		_product.GET("/getAllCategories", GetAllCategoriesHandler)
		_product.GET("/listProducts", ListProductsHandler)
		_product.POST("/createProduct", CreateProductHandler)
		_product.POST("/deleteProduct", DeleteProductHandler)
		// _product.POST("/updateProduct", UpdateProductHandler)
		api.GET("/hello", HelloWorldHandler)
	}

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		e.Logger.Errorf("Failed to marshal routes: %v", err)
	}
	log.Printf("Routes: %s", string(data))
}
