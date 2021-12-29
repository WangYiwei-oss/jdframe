package models

type UserRole struct {
	Id     int `gorm:"column:id;primaryKey"`
	UserId int `gorm:"column:user_id"`
	RoleId int `gorm:"column:role_id"`
}
