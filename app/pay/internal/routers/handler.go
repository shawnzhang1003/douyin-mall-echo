package routers

import (
	"errors"
	"net/http"

	"github.com/MakiJOJO/douyin-mall-echo/app/pay/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/app/pay/utils"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"

	// "github.com/MakiJOJO/douyin-mall-echo/app/pay/internal/logic"

	"github.com/labstack/echo/v4"
)

func AliPayHandler(c echo.Context) error {
	orderid := c.FormValue("orderid")
	if orderid == "" {
		return errors.New("orderid is empty")
	}
	totalPrice := c.FormValue("totalPrice")
	if totalPrice == "" {
		return errors.New("totalPrice is empty")
	}
	pay := utils.ZfbPay(orderid, totalPrice)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"payUrl": pay,
	})
	return c.String(http.StatusOK, pay)
}

func HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}
	c.Logger().Info("Hello World")
	mtl.Logger.Info("Hello World")

	return c.JSON(http.StatusOK, resp)
}

func healthHandler(c echo.Context) error {
	if dal.DbInstance == nil {
		return c.JSON(http.StatusInternalServerError, "database is not connected")
	}
	return c.JSON(http.StatusOK, dal.DbInstance.Health())
}
