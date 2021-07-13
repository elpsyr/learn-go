package fmt

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	file, err := os.OpenFile("./trace.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	Logger = log.New(io.MultiWriter(os.Stdout, file), "My Log: ", log.LstdFlags)
}
