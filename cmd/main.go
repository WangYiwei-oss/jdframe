package main

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/src/middlewares"
	"github.com/WangYiwei-oss/jdframe/src/testcontroller"
)

func main() {
	jdft.GetLogger("Global").Info("这是一个测试")
	jdft.NewJdft().DefaultBean().
		Attach(middlewares.RBAC()).
		Mount("v1", testcontroller.NewUserController()).
		Mount("v2", testcontroller.NewUserController()).
		CronTask("0/3 * * * * *", func() {}).
		Launch()
}
