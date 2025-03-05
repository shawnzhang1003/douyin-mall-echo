package routers

import (
	"fmt"
	"net/http"

	"github.com/MakiJOJO/douyin-mall-echo/app/auth/internal/dal"
	"github.com/MakiJOJO/douyin-mall-echo/app/auth/internal/logic"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	"github.com/golang-jwt/jwt/v4"
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
		return echo.NewHTTPError(http.StatusInternalServerError, "database is not connected")
	}
	return c.JSON(http.StatusOK, dal.DbInstance.Health())
}

func RefreshTokenHandler(c echo.Context) error {
	// redirect := c.QueryParam("redirect")
	// 将refresh token存储在cookie session中, 为了防止xss和csrf攻击, 需要设置cookie的HttpOnly和SameSite属性
	sess, err := session.Get("dymallsess", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "not found session, error: ", err)
	}
	refreshToken := sess.Values["RefreshToken"]
	if refreshToken == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "refresh token not found")
	}
	access, refresh, err := logic.JWTAuthService().RefreshToken(c.Request().Context(), refreshToken.(string))
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {
			return echo.NewHTTPError(http.StatusUnauthorized, "refresh token expired.")
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "refresh token invalid error: ", err)
	}
	sess.Options = &sessions.Options{
		Path:     "/api/v1/auth/refresh",
		MaxAge:   86400 * 30, //单位秒, 30天
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		// Secure:   true,
	}
	sess.Values["RefreshToken"] = refresh
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "save session error: ", err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": fmt.Sprintf("Bearer %v", access),
	})
}
