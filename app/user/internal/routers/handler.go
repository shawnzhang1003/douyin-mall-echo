package routers

import (
	"net/http"

	"github.com/MakiJOJO/douyin-mall-echo/app/user/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/app/user/internal/logic"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/user"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

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

type LoginReq struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func LoginHandler(c echo.Context) error {
	req := &LoginReq{}
	if err := c.Bind(req); err != nil {
		c.Logger().Errorf("bind error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "login using invalid form")
	}
	u, token, err := logic.UserLogin(req.Email, req.Password)
	if err != nil {
		c.Logger().Errorf("user login error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "email or password not correct")
	}
	sess, err := session.Get("dymallsess", c)
	if err != nil {
		c.Logger().Errorf("session get error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "get session error: %v")
	}
	sess.Options = &sessions.Options{
		Path:     "/api/v1/auth/refresh",
		MaxAge:   86400 * 30, //单位秒, 30天
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		// Secure:   true,
	}
	sess.Values["RefreshToken"] = token.RefreshToken
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		c.Logger().Errorf("session save error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "save session error: %v")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":     int32(u.ID),
		"accessToken": "Bearer " + token.AccessToken,
		// "refreshToken": token.RefreshToken,
	})
}

type RegisterReq struct {
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

func RegisterHandler(c echo.Context) error {
	req := &RegisterReq{}
	if err := c.Bind(req); err != nil {
		c.Logger().Errorf("bind req error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "register using invalid form")
	}
	u, err := logic.UserRegister(req.Email, req.Password, req.ConfirmPassword)
	if err != nil {
		c.Logger().Errorf("user register failed: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "user register failed, invalid email or password")
	}
	c.JSON(http.StatusOK, user.RegisterResp{
		UserId: int32(u.ID),
	})
	return nil
}

// handleLogout 处理用户登出请求
func LogoutHandler(c echo.Context) error {
	// 获取 session
	sess, err := session.Get("dymallsess", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "无法获取 session")
	}

	// 清除所有 session 数据
	sess.Options = &sessions.Options{
		Path:     "/api/v1/auth/refresh",
		MaxAge:   -1, // 设置 cookie 立即过期
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		// Secure:   true,
	}
	// sess.Values = make(map[interface{}]interface{}) // 清空 session 数据

	// 保存更改
	// err = sess.Save(c.Request(), c.Response())
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, "保存 session 失败")
	// }
	// 记录用户登出/修改密码的时间(now()), 以便分析用户活跃度和让其他浏览器中Cookie中的session失效
	if err := logic.UserLogout(c.Get("userid").(uint)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "用户登出失败")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "登出成功",
	})
}

func GetUserInfoHandler(c echo.Context) error {
	// 获取用户信息
	userID := c.Get("userid").(uint)
	u, err := logic.GetUserInfo(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "获取用户信息失败")
	}
	return c.JSON(http.StatusOK, u)
}
