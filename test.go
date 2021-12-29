package main

import (
	"context"
	"fmt"
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
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	fmt.Println(run(ctx))
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	fmt.Println(run(ctx))
}
