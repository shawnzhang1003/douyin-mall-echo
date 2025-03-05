package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	msgErrorTokenExpiredAndNoRefreshToken = 1001
)

type IsWhiteListPath func(path string) bool

type Claims struct {
	UserID   uint   `json:"userid"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func JWTAuthMiddleware(secretKey string, opt IsWhiteListPath) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 从配置中获取白名单路径
			if opt(c.Request().URL.Path) {
				return next(c)
			}

			// 获取 Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, `error: 未提供认证信息`)
			}

			// 检查 Bearer token
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				return echo.NewHTTPError(http.StatusUnauthorized, `error: 认证格式错误`)
			}

			// 解析 token
			claims := &Claims{}
			token, err := jwt.ParseWithClaims(parts[1], claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil // 这里的密钥应该从本服务的配置文件中读取,所以作为参数传进来
			}, jwt.WithValidMethods([]string{"HS256"}))

			// 检查是否是因为 token 过期导致的错误
			if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors == jwt.ValidationErrorExpired {
				// 重定向到认证服务刷新续期 token的api
				return c.Redirect(http.StatusFound, "/api/v1/auth/refresh")
			}

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, `error: 无效的token`, err.Error())

			}
			// 将用户信息存储到上下文中
			// 后续的handler可以通过 c.Get("userid") c.Get("username") 获取用户信息
			c.Set("userid", claims.UserID)
			c.Set("username", claims.Username)
			return next(c)
		}
	}
}
