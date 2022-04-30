package models

import (
	"fmt"
)

// User 默认角色表，想使用默认角色表请继承
type User struct {
	ID       int64  `gorm:"primarykey;column:id" json:"id"`
	UserName string `gorm:"column:user_name;type:varchar(25);not null;" json:"user_name"`
	Password string `gorm:"column:password;type:varchar(25);not null;" json:"password"`
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
	return fmt.Sprintf("%s-%s", u.ID, u.UserName)
}
