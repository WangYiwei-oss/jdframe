package main

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/src/logger"
	_ "github.com/WangYiwei-oss/jdframe/src/logger"
	"github.com/WangYiwei-oss/jdframe/src/qrcode"
	"github.com/WangYiwei-oss/jdframe/src/testcontroller"
)

func main(){
	logger.GetLogger("Global").Info("这是一个测试")
	jdft.NewJdft().
		Beans(jdft.NewGormAdapter(),qrcode.NewQrCode()).
		Mount("v1",testcontroller.NewUserController()).
		Mount("v2",testcontroller.NewUserController()).
		CronTask("0/3 * * * * *", func() {}).
		Launch()
}
