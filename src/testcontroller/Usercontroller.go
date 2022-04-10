package testcontroller

import (
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/configs"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/src/models"
	"github.com/gin-gonic/gin"
	"time"
)

type User struct {
	Id       int    `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}
type UserController struct {
	DB    *configs.GormAdapter        //`inject:"-"`
	Redis *configs.Jedis              `inject:"-"`
	Snow  *configs.SnowflakeGenerator `inject:"-"`
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) GetUserName(ctx *gin.Context) (int, string) {
	//user := User{}
	//u.DB.First(&user)
	//passwd := ctx.Query("password")
	//fmt.Println(passwd)
	//fmt.Println(user)
	return 1, ""
}

func (u *UserController) GetUserName3(ctx *gin.Context) (int, jdft.Json) {
	ret := []interface{}{
		"asd", 111, "zxcz",
	}
	return 1, ret
}

func (u *UserController) GetUserName2(ctx *gin.Context) (int, jdft.Json) {
	ret := models.NewUserModel()
	ret.UserName = ctx.PostForm("username")
	jdft.GetLogger("Global").Info("执行了用户controller")
	jdft.DoTask(u.TestTask, u.CallBack)
	return 1, ret
}

func (u *UserController) TestTask(params ...interface{}) {
	time.Sleep(time.Second * 3)
	fmt.Println("这是一个Task")
}

func (u *UserController) CallBack() {
	fmt.Println("执行了CallBack")
}

func (u *UserController) Build(jdft *jdft.Jdft) {
	jdft.Handle("GET", "login", u.GetUserName)
	jdft.Handle("POST", "login2", u.GetUserName2)
	jdft.Handle("GET", "login3", u.GetUserName3)
}
