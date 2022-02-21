package models

import (
	"fmt"
)

// User 默认角色表，像使用默认角色表请继承
type User struct {
	ID       uint   `gorm:"primarykey"`
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
