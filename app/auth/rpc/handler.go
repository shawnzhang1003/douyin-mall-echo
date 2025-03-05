package rpc

import (
	"context"
	"errors"
	"log"

	"github.com/MakiJOJO/douyin-mall-echo/app/auth/internal/logic"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v4"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

var _ auth.AuthService = &AuthServiceImpl{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	if req.GetUserId() <= 0 {
		err = errors.New("invalid user id")
		return
	}
	if req.GetUsername() == "" {
		err = errors.New("invalid username")
		return
	}
	accessToken, refreshToken, err := logic.JWTAuthService().GenerateToken(uint(req.UserId), req.GetUsername())
	if err != nil {
		return nil, err
	}
	resp = &auth.DeliveryResp{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}
	return
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	var reqtoken = req.GetToken()
	if reqtoken == "" {
		err = errors.New("blank token")
		return
	}
	token, err := logic.JWTAuthService().ValidateToken(reqtoken)
	if token.Valid {
		resp = &auth.VerifyResp{
			Res: true,
		}
		claims := token.Claims.(jwt.MapClaims)
		log.Println("token verify success")
		log.Println(claims)
	} else {
		resp = &auth.VerifyResp{
			Res: true,
		}
		log.Println(err)
	}
	return
}

func (s *AuthServiceImpl) RefreshTokenByRPC(ctx context.Context, req *auth.RefreshTokenReq) (resp *auth.RefreshResp, err error) {
	access, refresh, err := logic.JWTAuthService().RefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}
	resp = &auth.RefreshResp{
		Token:        access,
		RefreshToken: refresh,
	}
	return
}
