package testcontroller

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/gin-gonic/gin"
)

type UserController struct {

}

func NewUserController() *UserController {
	return &UserController{}
}

func GetUserName(ctx *gin.Context)(int,string){
	return -400,"wyw"
}

func (u *UserController)Build(jdft *jdft.Jdft){
	jdft.Handle("GET","login",GetUserName)
}
