package models

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	Id     int `gorm:"column:id;primaryKey"`
	UserId int `gorm:"column:user_id"`
	RoleId int `gorm:"column:role_id"`
}

func (u *UserRole) TableName() string {
	return "user_roles"
}
