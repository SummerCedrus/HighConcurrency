package worker

import (
	"job"
	"sync"
	"fmt"
)
const MAXWORKCNT = 100
var Workers WorkerPool
type Worker struct{
	JobChan job.JobChan
}

type WorkerPool struct {
	JobChans chan job.JobChan
	quit chan bool
}
func (worker *Worker) Run(){
	worker.JobChan = make(chan job.Job)

	go func() {
		for{
			//将空闲worker的jobChan放回池子，好接收新任务
			Workers.JobChans <- worker.JobChan
			select {
			case job := <- worker.JobChan:
				job.Do()
			}
		}
	}()
}

func InitWorkPool() {
	Workers.JobChans = make(chan job.JobChan, MAXWORKCNT)
	Workers.quit = make(chan bool)
	for i := 0; i < MAXWORKCNT; i++{
		wolker := new(Worker)
		wolker.Run()
	}
}

func Dispatcher(wg sync.WaitGroup) {
	InitWorkPool()
	for {
		select {
		//把新任务给池子里的worker的jobchan
		case job, ok := <-job.JobQueue:
			if !ok{
				fmt.Println(job.Content)
			}
			go func() {
				jobChan := <-Workers.JobChans
				jobChan <- job
			}()
		case quit := <-Workers.quit:
			if quit {
				close(Workers.JobChans)
				close(job.JobQueue)
				wg.Done()
				return
			}
		}

	}
}