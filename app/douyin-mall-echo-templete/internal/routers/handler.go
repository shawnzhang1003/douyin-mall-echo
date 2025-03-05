package routers

import (
	"net/http"
	"strconv"

	"github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/internal/logic"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	"github.com/labstack/echo/v4"
)

type hellowReq struct {
	Name string `json:"name" form:"name"` // 这样就能接受json和formData格式的参数
}

// 示例代码
func HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}
	var req hellowReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bind request failed")
	}
	resp["name"] = req.Name
	// 第二种获取username和userid的方法 c.Get("userid") c.Get("username")
	// 但是前提需要在该handler前经过jwt中间件鉴权获取用户信息才有值
	resp["userID"] = strconv.Itoa(int(c.Get("userid").(uint)))

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

func LoginHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	logic.Login(username, password)
	return nil
}

func RegisterHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	newUser, err := logic.Register(username, password)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, *newUser)
	return nil
}
