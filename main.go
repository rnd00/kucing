package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/rnd00/kucing/cache"
	"github.com/rnd00/kucing/compare"
	"github.com/rnd00/kucing/download"
	"github.com/rnd00/kucing/helpers"
)

const (
	cacheDuration time.Duration = time.Second * 2
	workerAmount  int           = 2
	jobsAmount    int           = 12
)

func main() {
	DownloadWorker(workerAmount, jobsAmount)
}

func DownloadWorker(workerAmt, jobsAmt int) {
	var wg sync.WaitGroup
	jobs := make(chan int, jobsAmt)
	cache := cache.NewCache().AddTimeout(cacheDuration)

	// open up subthread
	for i := 0; i < workerAmt; i++ {
		wg.Add(1)
		logs := make(chan *helpers.Log)
		// spin up another thread
		go worker(i+1, jobs, logs, cache)
		go logger(i+1, logs, &wg)
	}

	// assign some jobs
	for i := 0; i < jobsAmt; i++ {
		jobs <- i + 1
		// add sleep so that it has a delay
		time.Sleep(time.Second)
	}
	// close when done
	close(jobs)

	// put waitgroup to wait
	wg.Wait()
	procEndLog := helpers.NewLogInfo("[MAIN] All the channels are closed and workers are done.")
	procEndLog.Print()
}

func worker(workerTag int, jobsC <-chan int, logsC chan<- *helpers.Log, cache *cache.Cac) {
	defer close(logsC)

	amtOfJobs := 0

	// while jobsC is not closed
	for n := range jobsC {
		amtOfJobs++

		// start log
		startLogMsg := fmt.Sprintf("[WORKER %d] job %d\t| Start", workerTag, n)
		startL := helpers.NewLogInfo(startLogMsg)
		logsC <- startL

		// make another loop here
		// if fail then retry downloading
		jobDone := false
		attempt := 0

		for !jobDone {
			// retry attempt log
			attempt += 1
			retryAttemptLog := fmt.Sprintf("[WORKER %d] job %d\t| Attempt %d", workerTag, n, attempt)
			retryL := helpers.NewLogInfo(retryAttemptLog)
			logsC <- retryL

			// get data
			bytedata, err := download.Cat()
			if err != nil {
				errL := helpers.NewLogError(*helpers.NewError(err))
				logsC <- errL
				// retry this loop by skipping
				continue
			}

			key := compare.MakeKey(helpers.ToBase64Bytes(bytedata))
			debugLog := fmt.Sprintf("[WORKER %d] job %d\t| generated key: %s", workerTag, n, key)
			logsC <- helpers.NewLogDebug(debugLog)

			exist := cache.CheckKey(key)
			if exist {
				errL := helpers.NewLogError(*helpers.NewError(errors.New("Same generated key existed in the cache")))
				logsC <- errL
				// retry this loop by skipping
				continue
			}

			cache.SetKeyWithTimeout(key)
			var filename string
			if filename, err = helpers.WriteToFile(bytedata); err != nil {
				errL := helpers.NewLogError(*helpers.NewError(err))
				logsC <- errL
				continue
			}
			fileCreatedL := fmt.Sprintf("[WORKER %d] job %d\t| File written: %s", workerTag, n, filename)
			fileL := helpers.NewLogInfo(fileCreatedL)
			logsC <- fileL

			jobDone = true
		}

		jobDoneMsg := fmt.Sprintf("[WORKER %d] job %d\t| done, moving to the next one", workerTag, n)
		jobDoneL := helpers.NewLogInfo(jobDoneMsg)
		logsC <- jobDoneL
	}

	stopLogMsg := fmt.Sprintf("[WORKER %d] %d jobs done, closing", workerTag, amtOfJobs)
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
