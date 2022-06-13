package main

import (
	"fmt"
	"sync"

	"github.com/rnd00/kucing/download"
	"github.com/rnd00/kucing/helpers"
)

func main() {
	// testingDownload()
	testingDownloadUsingWorker()
}

func testingDownloadUsingWorker() {
	workerAmt := 2
	jobsAmt := 12

	var wg sync.WaitGroup
	jobs := make(chan int, jobsAmt)

	// open up subthread
	for i := 0; i < workerAmt; i++ {
		wg.Add(1)
		logs := make(chan *helpers.Log)
		go worker(i+1, jobs, logs)
		go logger(i+1, logs, &wg)
	}

	// assign some jobs
	for i := 0; i < jobsAmt; i++ {
		jobs <- i + 1
	}
	// close when done
	close(jobs)

	// put waitgroup to wait
	wg.Wait()
	procEndLog := helpers.NewLogInfo("[MAIN] All the channels are closed and workers are done.")
	procEndLog.Print()
}

func worker(workerTag int, jobsC <-chan int, logsC chan<- *helpers.Log) {
	defer close(logsC)

	// while jobsC is not closed
	for n := range jobsC {
		// start log
		startLogMsg := fmt.Sprintf("[WORKER %d] job %d\t| Start", workerTag, n)
		startL := helpers.NewLogInfo(startLogMsg)
		logsC <- startL

		// get data
		bytedata, err := download.Cat()
		if err != nil {
			errL := helpers.NewLogError(*helpers.NewError(err))
			logsC <- errL

			// skip this loop
			continue
		}
		// key := compare.MakeKey(bytedata)
		// exist := cache.CheckExistence(key)
		// if exist { log <- errors.New("data already exist"); return }
		// helpers.WriteToFile(bytedata)

		// end log (debug)
		endLogMsg := fmt.Sprintf("[WORKER %d] job %d\t| Retrieved with length %d", workerTag, n, len(bytedata))
		endL := helpers.NewLogInfo(endLogMsg)
		logsC <- endL
	}

	stopLogMsg := fmt.Sprintf("[WORKER %d] jobs done, closing", workerTag)
	stopL := helpers.NewLogInfo(stopLogMsg)
	logsC <- stopL
	return
}

func logger(workerTag int, c chan *helpers.Log, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range c {
		fmt.Println(data)
	}
	stopLoggerMsg := fmt.Sprintf("[LOGGER %d] worker is done, closing", workerTag)
	fmt.Println(helpers.NewLogInfo(stopLoggerMsg))
}
