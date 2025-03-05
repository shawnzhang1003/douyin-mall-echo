package routers

import (
	"net/http"
	"strconv"

	"github.com/MakiJOJO/douyin-mall-echo/app/cart/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/app/cart/internal/logic"
	"github.com/MakiJOJO/douyin-mall-echo/app/cart/model"
	"github.com/labstack/echo/v4"
)

func HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func healthHandler(c echo.Context) error {
	if dal.DbInstance == nil {
		return c.JSON(http.StatusInternalServerError, "database is not connected")
	}
	return c.JSON(http.StatusOK, dal.DbInstance.Health())
}

func AddItemHandler(c echo.Context) error {
	item := &model.Cart{}
	if err := c.Bind(item); err != nil {
		return err
	}
	
	if err := logic.AddItem(item); err != nil {
		return err
	}

	resp := map[string]string{
		"message": "add item success!",
	}
	return c.JSON(http.StatusOK, resp)
}

func GetCartHandler(c echo.Context) error {
	userid, err := strconv.Atoi(c.QueryParam("userid"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	itemList, err := logic.GetCartReturnInfo(uint32(userid))
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, itemList)

	return nil
}

func EmptyCartHandler(c echo.Context) error {
	item := &model.Cart{}
	if err := c.Bind(item); err != nil {
		return err
	}

	if err := logic.EmptyCart(item.UserId); err != nil {
		return err
	}

	resp := map[string]string{
		"message": "empty cart success!",
	}
	return c.JSON(http.StatusOK, resp)
}
