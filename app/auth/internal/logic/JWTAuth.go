package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/MakiJOJO/douyin-mall-echo/app/auth/config"
	"github.com/MakiJOJO/douyin-mall-echo/app/auth/rpc/client"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/user"
	"github.com/golang-jwt/jwt/v4"
)

// jwt service
type JWTService interface {
	GenerateToken(userID uint, username string) (accessToken, refreshToken string, err error)
	ValidateToken(token string) (*jwt.Token, error)
	RefreshToken(ctx context.Context, refreshToken string) (access, refresh string, err error)
}
type authCustomClaims struct {
	UserID   uint   `json:"userid"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type jwtServices struct {
	secretKey        string
	refreshSecretKey string
	issuer           string
}

var _ JWTService = &jwtServices{}

// auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey:        getSecretKey(),
		refreshSecretKey: getRefreshSecretKey(),
		issuer:           "douyin-mall",
	}
}

func getSecretKey() string {
	secret := config.GlobalConfig.JWT.SecretKey
	if secret == "" {
		secret = "secret"
	}
	return secret
}
func getRefreshSecretKey() string {
	secret := config.GlobalConfig.JWT.RefreshSecretKey
	if secret == "" {
		secret = "refresh_secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(userID uint, username string) (accessToken, refreshToken string, err error) {
	refreshClaims := &authCustomClaims{
		userID,
		username,
		jwt.RegisteredClaims{
			//get expired after 30 Days
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Hour * 24)), // 30天有效期
			Issuer:    service.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	// 生成 access token
	accessClaims := &authCustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // 15分钟有效期
			Issuer:    service.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	//encoded string
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).
		SignedString([]byte(service.refreshSecretKey))
	if err != nil {
		return "", "", err
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).
		SignedString([]byte(service.secretKey))
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	// []string{"HS256"}是添加加密算法名称白名单,为什么要添加参考资料https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/

}

func (service *jwtServices) RefreshToken(ctx context.Context, refreshToken string) (access, refresh string, err error) {
	// token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
	// 	if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
	// 		return nil, fmt.Errorf("invalid token %s", token.Header["alg"])

	// 	}
	// 	return []byte(service.refreshSecretKey), nil
	// }, jwt.WithValidMethods([]string{"HS256"}))
	var claims authCustomClaims
	token, err := jwt.ParseWithClaims(refreshToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])

		}
		return []byte(service.refreshSecretKey), nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if token.Valid {
		// claims := token.Claims.(jwt.MapClaims)
		// userID := claims["userid"].(float64)
		// logoutTime := claims["logout"]
		// username := claims["username"].(string)
		userID := claims.UserID
		username := claims.Username
		// 得到用户登出/修改密码的时间, 如果refreshToken的签发时间在用户登出/修改密码之前, 则返回token过期
		// 这样可以防止用户登出后, 仍然可以在其他设备上使用refreshToken刷新token
		req := &user.LogoutTimeReq{
			UserId: int32(userID),
		}
		resp, err := client.UserClient.GetUserLogoutTime(ctx, req)
		if err != nil {
			return "", "", err
		}
		if claims.IssuedAt.Time.Before(resp.LogoutTime.AsTime()) {
			return "", "", fmt.Errorf("token is expired")
		}

		access, refresh, err = service.GenerateToken(uint(userID), username)
		if err != nil {
			return "", "", err
		}
	}
	return
}

// middleware
// func AuthorizeJWT() echo.HandlerFunc {
// 	return func(c echo.Context) {
// 		const BEARER_SCHEMA = "Bearer " //format is "Bearer xxxxxxxx"
// 		authHeader := c.GetHeader("Authorization")
// 		log.Println("authHeader is", authHeader)
// 		tokenString := authHeader[len(BEARER_SCHEMA):]
// 		log.Println("token is", tokenString)
// 		token, err := JWTAuthService().ValidateToken(tokenString)
// 		if token.Valid {
// 			claims := token.Claims.(jwt.MapClaims)
// 			log.Println("token verify success")
// 			log.Println(claims)
// 		} else {
// 			//todo, recreate token?
// 			log.Println(err)
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}

// 	}
// }
