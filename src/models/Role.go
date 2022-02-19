package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Role struct { //对应roles这张表，这张表定义了g的规则
	gorm.Model
	RoleId      int    `gorm:"column:role_id;primaryKey"`
	RoleName    string `gorm:"column:role_name"`
	RolePid     int    `gorm:"column:role_pid"`
	RoleComment string `gorm:"column:role_comment"`
}

func (r *Role) TableName() string {
	return "roles"
}

func (r *Role) String() string {
	return fmt.Sprintf("ID:%d 角色名:%s", r.RoleId, r.RoleName)
}
