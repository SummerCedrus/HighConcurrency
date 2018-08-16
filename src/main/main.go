package main

import (
	"job"

	"time"

	"sync"
	"worker"
)

func main(){
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go job.GenJobs(1000, time.Second, *wg)
	go worker.Dispatcher(*wg)
	wg.Wait()
	close(job.JobQueue)
}
