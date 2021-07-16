package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("finish doing something")
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func main() {
	//创建空context的两种方法
	ctx := context.Background() //返回一个空的context，不能被cancel，kv为空
	fmt.Println(ctx)

	//todoCtx := context.TODO()

	//context附魔
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	doSomething(ctx)

}
