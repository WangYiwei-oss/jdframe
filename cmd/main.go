package main

import (
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/WangYiwei-oss/jdframe/src/testcontroller"
)

func main(){
	jdft.NewJdft().
		Beans(jdft.NewGormAdapter()).
		Mount("v1",testcontroller.NewUserController()).
		Mount("v2",testcontroller.NewUserController()).
	Launch()
}
