package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MakiJOJO/douyin-mall-echo/app/user/model"
	"github.com/MakiJOJO/douyin-mall-echo/app/user/rpc/client"
	"github.com/MakiJOJO/douyin-mall-echo/app/user/util"
	"github.com/MakiJOJO/douyin-mall-echo/rpc_gen/kitex_gen/auth"
)

var ErrEncryptPasswordFail = errors.New("cannot encrypt password")

type Token struct {
	AccessToken  string
	RefreshToken string
}

func UserRegister(email, password, confirm_password string) (*model.User, error) {
	if password != confirm_password {
		return nil, errors.New("confirmed password does not match")
	}
	if !util.IsValidEmail(email) {
		return nil, fmt.Errorf("invalid email: %v", email)
	}
	password, err := util.Encrypt(password)
	if err != nil {
		return nil, ErrEncryptPasswordFail
	}
	u := &model.User{
		Email:    email,
		Password: password,
	}
	if err := model.AddUser(u); err != nil {
		return nil, err
	}
	return u, nil
}

func UserLogin(email, password string) (*model.User, *Token, error) {
	u, err := model.GetUserByEmail(email)
	if err != nil {
		return nil, nil, fmt.Errorf("get email from database error: %v", err)
	}
	if err := util.ValidatePassword(password, u.Password); err != nil {
		return nil, nil, errors.New("invalid password")
	}
	req := &auth.DeliverTokenReq{
		UserId:   int32(u.ID),
		Username: u.Email,
	}
	resp, err := client.AuthClient.DeliverTokenByRPC(context.Background(), req)
	if err != nil {
		return nil, nil, fmt.Errorf("get token from auth service error: %v", err)
	}

	return u, &Token{resp.Token, resp.RefreshToken}, nil
}

func UserLogout(userID uint) error {
	return model.UpdateUserLogoutTime(userID)
}

func GetUserLogoutTime(userID uint) (time.Time, error) {
	return model.GetUserLogoutTime(userID)
}
