package routers

import (
	"net/http"

	"github.com/MakiJOJO/douyin-mall-echo/app/ai_agent/internal/logic"
	"github.com/labstack/echo/v4"
)

func HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

type Query struct {
    UserId uint32 `json:"user_id" binding:"required"`
    Query  string `json:"query" binding:"required"`
}

func Ai_agentHandler(c echo.Context) error {
	var query Query
	err := c.Bind(&query)
	if err != nil {
		return err
	}
	if query.UserId == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "user_id is null, or you should use user_id")
	}

	if query.Query == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "query is null")
	}

	resp, err := logic.AgentInvoke(c.Request().Context() ,query.UserId, []byte(query.Query))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
