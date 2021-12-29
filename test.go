package main

import (
	"context"
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"time"
)

func job() chan string {
	ret := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		ret <- "success"
	}()
	return ret
}

func run(ctx context.Context) string {
	ret := job()
	select {
	case r := <-ret:
		return r
	case <-ctx.Done():
		return "killed"
	}
}

func main() {
	uu := jdft.UserRole{}
	jdft.NewGormAdapter().DB.Where("id=1").First(&uu)
	fmt.Println(uu)
}
