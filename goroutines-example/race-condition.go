package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var mtx sync.Mutex
var counter int32
var wait sync.WaitGroup

var ch = make(chan int, 1)

func UnsafeIncCounter() {
	defer wait.Done()
	for i := 0; i < 10000; i++ {
		mtx.Lock()
		counter++
		mtx.Unlock()
	}
}

func AtomicIncCounter() {
	defer wait.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddInt32(&counter, 1)
	}
}

func ChannelIncCounter() {
	defer wait.Done()
	for i := 0; i < 10000; i++ {
		int := <-ch
		int++
		ch <- int
	}
}

func main() {
	wait.Add(2)
	//go UnsafeIncCounter()
	//go UnsafeIncCounter()

	//go AtomicIncCounter()
	//go AtomicIncCounter()

	go ChannelIncCounter()
	go ChannelIncCounter()
	ch <- 0

	wait.Wait()
	fmt.Println(<-ch)
	//fmt.Println(counter)
}
