package rpc

import (
	"context"
	"errors"

	"github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/internal/logic"
	user "github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	username := req.GetEmail()
	password := req.GetPassword()
	confirmPassword := req.GetConfirmPassword()
	if password != confirmPassword {
		err = errors.New("password must be the same as ConfirmPassword")
		return
	}

	newUser, err := logic.Register(username, password)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	return
}

func (s *UserServiceImpl) GetUserLogoutTime(ctx context.Context, req *user.LogoutTimeReq) (resp *user.LogoutTimeResp, err error) {

	return
}
