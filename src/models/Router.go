package models

type Router struct {
	RouterId     int    `gorm:"column:r_id;primaryKey"`
	RouterName   string `gorm:"column:r_name"`
	RouterUri    string `gorm:"column:r_uri"`
	RouterMethod string `gorm:"column:r_method"`
	RouterStatus int    `gorm:"column:r_status"`
	RoleName     string `gorm:"column:role_name"`
}

func (r *Router) TableName() string {
	return "routers"
}
