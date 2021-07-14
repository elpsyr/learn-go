package main

import (
	"fmt"
	"learn-go/concurrent-models-example/runner"
	"time"
)

func createTask() func(int) {
	return func(id int) {
		time.Sleep(time.Second)
		fmt.Printf("Task Complete # %d \n", id)

	}
}

func main() {
	r := runner.New(5 * time.Second)

	r.AddTask(createTask(), createTask(), createTask())
	err := r.Start() //开始后主线程开始等待返回结果

	switch err {
	case runner.ErrInterrupt:
		fmt.Println("tasks interrupted")
	case runner.ErrTimeout:
		fmt.Println("tasks timeout")
	default:
		fmt.Println("all tasks finished")
	}

}
