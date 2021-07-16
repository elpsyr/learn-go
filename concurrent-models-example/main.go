package main

import (
	"fmt"
	"io"
	"learn-go/concurrent-models-example/pool"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// DBConnection	定义的一个资源
type DBConnection struct {
	id int32
}

func (D DBConnection) Close() error {
	fmt.Println("database closed , #" + fmt.Sprint(D.id))
	return nil
}

var counter int32

func Factory() (io.Closer, error) {
	atomic.AddInt32(&counter, 1)
	return DBConnection{
		id: counter,
	}, nil
}

func performQuery(query int, pool *pool.Pool) {
	defer wg.Done()

	resource, err := pool.AcquireResource()

	if err != nil {
		log.Fatalln(err)
	}
	defer pool.ReleaseResource(resource)

	//t := rand.Int()%10 + 1
	//time.Sleep(time.Duration(t) * time.Second)
	time.Sleep(time.Second)
	fmt.Println("finish query" + fmt.Sprint(query))

}

var wg sync.WaitGroup

func main() {
	p, err := pool.New(Factory, 5)
	if err != nil {
		log.Fatalln(err)
	}

	num := 10
	wg.Add(num)
	for id := 0; id < 5; id++ {

		go performQuery(id, p)
	}
	time.Sleep(2 * time.Second)
	for id := 5; id < num; id++ {

		go performQuery(id, p)
	}

	wg.Wait()

	p.Close()

	fmt.Println("pool model run ends")
}

//func createTask() func(int) {
//	return func(id int) {
//		time.Sleep(time.Second)
//		fmt.Printf("Task Complete # %d \n", id)
//
//	}
//}
//
////func main() {
////	r := runner.New(5 * time.Second)
////
////	r.AddTask(createTask(), createTask(), createTask())
////	err := r.Start() //开始后主线程开始等待返回结果
////
////	switch err {
////	case runner.ErrInterrupt:
////		fmt.Println("tasks interrupted")
////	case runner.ErrTimeout:
////		fmt.Println("tasks timeout")
////	default:
////		fmt.Println("all tasks finished")
////	}
////
////}
