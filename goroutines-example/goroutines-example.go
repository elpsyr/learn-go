package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func play(name string, ch chan int) {

	defer wg.Done() //写在前面
	for true {
		ball, ok := <-ch
		if !ok { //通道关闭
			fmt.Printf("%s win!\n", name)
			return //结束方法
		}

		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		n := r.Intn(100)
		if n%10 == 0 { //把球打飞
			close(ch)
			fmt.Printf("%s lose the ball\n", name)
			return
		}
		ball++
		fmt.Printf("%s receive the ball %dtimes \n", name, ball)
		ch <- ball
	}
}

//func main() {
//	ch := make(chan int, 0)
//	wg.Add(2)//等待2个任务完成
//	go play("one",ch)
//	go play("two",ch)
//	ch<-0//线程启动后才能启用channel，不然没有人接收
//	wg.Wait()//等待其他goroutines完成
//	fmt.Println("exit")
//}
