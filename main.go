package main

import (
	"fmt"
	"time"

	"github.com/rnd00/kucing/download"
)

func main() {
	testingDownload()
}

func testingDownload() {
	dataC := make(chan []byte)
	errC := make(chan error)

	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Second / 2)
			download.Routine(dataC, errC)
		}
	}()

	go func() {
		for data := range dataC {
			fmt.Println("[RETRIEVED]", time.Now().Format(time.RFC3339), len(data))
		}
	}()

	go func() {
		for err := range errC {
			fmt.Println("[ERROR]", time.Now().Format(time.RFC3339), err)
		}
	}()

	fmt.Scanln()
}
