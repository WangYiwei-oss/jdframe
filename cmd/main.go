package main

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/src/testcontroller"
	"log"
)

func main(){
	jdft.NewJdft().
		Beans(jdft.NewGormAdapter()).
		Mount("v1",testcontroller.NewUserController()).
		Mount("v2",testcontroller.NewUserController()).
		CronTask("0/3 * * * * *", func() {
		log.Println("执行定时任务")
	}).
	Launch()
}
