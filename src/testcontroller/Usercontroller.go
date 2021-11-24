package testcontroller

import (
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/gin-gonic/gin"
)
type User struct {
	Id int	`gorm:"column:id"`
	Username string	`gorm:"column:username"`
	Password string `gorm:"column:password"`
}
type UserController struct {
	DB *jdft.GormAdapter `inject:"-"`
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController)GetUserName(ctx *gin.Context)(int,string){
	user := User{}
	u.DB.First(&user)
	fmt.Println(user)
	return -400,"wyw"
}

func (u *UserController)Build(jdft *jdft.Jdft){
	jdft.Handle("GET","login",u.GetUserName)
}
