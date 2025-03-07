package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/MakiJOJO/douyin-mall-echo/app/user/internal/dal"
	"gorm.io/gorm"
)

var ErrDuplicateKey = errors.New("uid or email has existed")

type User struct {
	gorm.Model
	Email      string    `gorm:"type:varchar(50);unique;column:email"`
	Password   string    `gorm:"type:varchar(100);column:password"`
	UserName   string    `gorm:"type:varchar(50);column:user_name"`
	Gender     string    `gorm:"type:varchar(10);column:gender"`
	LogoutTime time.Time `gorm:"type:datetime;column:logout_time"`
}

func (u *User) TableName() string {
	return "users"
}

func AddUser(user *User) error {
	db := dal.DB
	if isEmailExists(user.Email) {
		return ErrDuplicateKey
	}
	if err := db.Create(user).Error; err != nil {
		return errors.New(fmt.Sprintf("add user failed: %v", err))
	}
	return nil
}

func isEmailExists(email string) bool {
	db := dal.DB
	var user User
	result := db.First(&user, "email = ?", email)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func GetUserByID(id uint) (*User, error) {
	db := dal.DB
	var user User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	db := dal.DB
	var user User
	if err := db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserLogoutTime(id uint) error {
	if err := dal.DB.Model(&User{}).Where("id = ?", id).Update("logout_time", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func GetUserLogoutTime(id uint) (time.Time, error) {
	var user User
	if err := dal.DB.First(&user, "id = ?", id).Error; err != nil {
		return time.Time{}, err
	}
	return user.LogoutTime, nil
}
