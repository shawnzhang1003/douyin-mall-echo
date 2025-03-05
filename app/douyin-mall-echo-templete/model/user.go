package model

import (
	"strings"

	"github.com/MakiJOJO/douyin-mall-echo/app/douyin-mall-echo-templete/internal/dal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	gorm.Model
	Username  string  `gorm:"type:varchar(100);unique" form:"username" json:"username"`
	Password  string  `json:"-" gorm:"not null"`
	FirstName string  `gorm:"type:varchar(100)" form:"firstname" json:"firstname"`
	LastName  string  `gorm:"type:varchar(100)" form:"lastname" json:"lastname"`
	Nickname  string  `gorm:"type:varchar(100)" form:"nickname" json:"nickname"`
	Avatar    string  `gorm:"type:varchar(10)" form:"avatar" json:"avatar"`
	Email     string  `gorm:"type:varchar(50)" form:"email" json:"email"`
	Area      string  `gorm:"type:varchar(100)" form:"area" json:"area"` // country
	Roles     []*Role `gorm:"many2many:user_roles;" form:"roles" json:"roles"`
	//Password  string `gorm:"type:varchar(100)" form:"password"`
	//todo, last login,  date joined
}
type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(30);unique" form:"name" json:"name"`
}

func (u *User) TableName() string {
	return "users"
}

func CreateUser(u *User) error {
	condition := User{}
	var role = &Role{}

	//lowercase username
	condition.Username = strings.ToLower(u.Username)
	assign := User{}
	//log.Println("CreateOrUpdateUser", u)
	if err := dal.DB.Preload(clause.Associations).Where(condition).Assign(assign).FirstOrCreate(u).Error; err != nil {
		return err
	}
	// FirstOrCreate函数不会自动更新关联表的数据，需要手动更新,即使使用了Assign
	if len(u.Roles) == 0 {
		if err := dal.DB.Where(Role{Name: "USER"}).FirstOrCreate(role).Error; err != nil {
			return err
		}
		u.Roles = append(u.Roles, role)
		if err := dal.DB.Save(u).Error; err != nil {
			return err
		}
	}
	return nil
}