package main

import (
	"fmt"
	"github.com/casbin/casbin"
)

func main() {
	sub := "admin"
	obj := "/depts"
	act := "POST"
	e := casbin.NewEnforcer("src/casbin/model.conf", "src/casbin/p.csv")
	ok := e.Enforce(sub, obj, act)
	if ok {
		fmt.Println("验证通过")
	} else {
		fmt.Println("验证不通过")
	}
}
