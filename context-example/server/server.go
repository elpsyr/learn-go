package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

	//获取context
	fmt.Println("handler start")
	ctx := r.Context()

	complete := make(chan struct{})
	go func() {
		//do something
		time.Sleep(5 * time.Second)
		complete <- struct{}{}
	}()

	select {
	//模拟Done的模式
	case <-complete:
		fmt.Println("finish doing something")
		_, err := fmt.Fprintln(w, "hello world")
		if err != nil {
			log.Fatalln(err)
		}
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("handler ends")

	//默认情况下只要收到了请求就得把工作做完
	//time.Sleep(10*time.Second)
	//fmt.Println("keep working")
	//_, err := fmt.Fprintln(w, "hello world")
	//if err != nil {
	//	log.Fatalln(err)
	//}

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
