package main

import (
	"fmt"
	"unsafe"
)

var str string

type s1 struct {
	a int8
	c int16
	d int32
	b int64
}

func main() {
	//myfmt.Logger.Println("hello")
	//
	//request := downloader.InfoRequest{
	//	Bvids: []string{"BV1Ff4y187q9", "BV18f4y187kT"},
	//}
	//response, err := downloader.BatchDownloadVideoInfo(request)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, info := range response.Infos {
	//
	//	myfmt.Logger.Printf("title: %s \n desc: %s\n", info.Data.Title, info.Data.Desc)
	//}

	a := s1{
		a: 0,
		b: 0,
		c: 0,
		d: 0,
	}

	ints := []int{1, 2}
	ints = append(ints, 3, 4, 5)
	fmt.Println(cap(ints))
	fmt.Printf("n 的类型 %T n2占中的字节数是 %d", a, unsafe.Sizeof(a))

}
