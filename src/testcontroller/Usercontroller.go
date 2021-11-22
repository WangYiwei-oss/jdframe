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

func (u *UserController)Build(jdft *jdft.Jdft){
	jdft.Router.GET("login", func(context *gin.Context) {
		context.String(200,"helloword")
	})
}
