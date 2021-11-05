package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	//普通访问，一直等5秒
	//resp, err := http.Get("http://localhost:8080")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer resp.Body.Close()
	//respBody, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(respBody))

	//传递context
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()
	//创建request
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080", nil)

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}
