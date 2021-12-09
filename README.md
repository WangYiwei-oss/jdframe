# jdframe

## controller及依赖注入写法示例

```go
type User struct {
   Id       int    `gorm:"column:id"`
   Username string `gorm:"column:username"`
   Password string `gorm:"column:password"`
}
type UserController struct {
   DB *jdft.GormAdapter `inject:"-"`	//依赖注入,还需要再main函数中加入bean注册依赖
}

func NewUserController() *UserController {	//必须1
	return &UserController{}
}

func (u *UserController) GetUserName(ctx *gin.Context) (int, string) {	//控制器函数必须传gin.Context
	user := User{}
	u.DB.First(&user)
	fmt.Println(user)
	return -400, "wyw"
}

func (u *UserController) GetUserName3(ctx *gin.Context) int {
	return -400
}

func (u *UserController) Build(jdft *jdft.Jdft) {	//必须写build
	jdft.Handle("GET", "login", u.GetUserName)
    jdft.Handle("GET", "login2", u.GetUserName2)
}
```

main函数

```go
jdft.NewJdft().
   Beans(jdft.NewGormAdapter(), qrcode.NewQrCode()).//注册依赖
   Mount("v1", testcontroller.NewUserController()).//挂载路由
   Mount("v2", testcontroller.NewUserController()).
   CronTask("0/3 * * * * *", func() {}).	//定时器函数
   Launch()
```

