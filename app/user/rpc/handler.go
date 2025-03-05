package rpc

import (
	"context"

	"github.com/MakiJOJO/douyin-mall-echo/app/user/internal/logic"
	user "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServiceImpl struct{}

func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterResp, error) {
	u, err := logic.UserRegister(req.GetEmail(), req.GetPassword(), req.GetConfirmPassword())
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{UserId: int32(u.ID)}, nil
}

func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error) {
	u, _, err := logic.UserLogin(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &user.LoginResp{UserId: int32(u.ID)}, nil
}

func (s *UserServiceImpl) GetUserLogoutTime(ctx context.Context, req *user.LogoutTimeReq) (resp *user.LogoutTimeResp, err error) {

	logoutTime, err := logic.GetUserLogoutTime(uint(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	resp = &user.LogoutTimeResp{
		LogoutTime: timestamppb.New(logoutTime),
	}
	return
}
