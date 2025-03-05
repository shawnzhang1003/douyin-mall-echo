package logic

import (
	"errors"
	"strings"

	"github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/model"
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
)

func Login(username, password string) error {
	return nil

}

func Register(username, password string) (*model.User, error) {

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		//c.JSON(http.StatusBadRequest, gin.H{"result": "Parameters can't be empty"})
		return nil, errors.New("parameters can't be empty")
	}
	// 对密码加盐后进行加密
	var p string
	// p = Encryption(password)
	var user = model.User{
		Username: username,
		Password: p,
	}
	model.CreateUser(&user)
	mtl.Logger.Info("Register success")

	return &user, nil
}
