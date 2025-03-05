package routers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
)

func RegisterRoutes(e *echo.Echo) {

	e.GET("/health", HelloWorldHandler)
	api := e.Group("/api/v1")
	{
		api.POST("/ai_agent", Ai_agentHandler)
		api.GET("/hello", HelloWorldHandler)
	}

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		e.Logger.Errorf("Failed to marshal routes: %v", err)
	}
	log.Printf("Routes: %s", string(data))

}
