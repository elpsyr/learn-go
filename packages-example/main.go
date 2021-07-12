package main

import (
	"learn-go/packages-example/downloader"
	myfmt "learn-go/packages-example/fmt"
)

func main() {
	myfmt.Logger.Println("hello")

	request := downloader.InfoRequest{
		Bvids: []string{"BV1Ff4y187q9", "BV18f4y187kT"},
	}
	response, err := downloader.BatchDownloadVideoInfo(request)
	if err != nil {
		panic(err)
	}

	for _, info := range response.Infos {

		myfmt.Logger.Printf("title: %s \n desc: %s\n", info.Data.Title, info.Data.Desc)
	}

}
