package models

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId   int    `gorm:"column:user_id;primaryKey;AUTO_INCREMENT"`
	UserName string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	RoleName string `gorm:"column:role_name"`
}

func (u *User) TableName() string {
	return "users"
}

func NewUserModel() *User {
	return &User{}
}

func (u *User) JModelName() string {
	return "UserModel"
}

func (u *User) String() string {
	return fmt.Sprintf("%s:%s", u.UserName, u.RoleName)
}
