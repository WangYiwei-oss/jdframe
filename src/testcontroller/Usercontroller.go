package testcontroller

import (
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/src/models"
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

func (u *UserController)GetUserName3(ctx *gin.Context)int{
	return -400
}

func (u *UserController)GetUserName2(ctx *gin.Context)(int,jdft.JModel){
	ret := models.NewUserModel()
	ret.UserName="wyw"
	ret.Password="123"
	return -400,ret
}

func (u *UserController)Build(jdft *jdft.Jdft){
	jdft.Handle("GET","login",u.GetUserName)
	jdft.Handle("GET","login2",u.GetUserName2)
	jdft.Handle("GET","login3",u.GetUserName3)
}
