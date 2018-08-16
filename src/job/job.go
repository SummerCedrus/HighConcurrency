package job

import (
	"time"
	"fmt"
	"sync"
)
const MAXQUEUELEN = 1000
var JobQueue chan Job
type JobChan chan Job

type Job struct {
	Content int
	CostSec time.Duration
}

func (job *Job) Do(){
	time.Sleep(job.CostSec)
	fmt.Println(job.Content)
}
func InitJobQueue(){
	JobQueue = make(chan Job, MAXQUEUELEN)
}
//
func GenJobs(jobCnt int, costTime time.Duration, wg sync.WaitGroup){
	InitJobQueue()
	for i := 0;i < jobCnt; i++{
		job := Job{CostSec:costTime, Content:i}

		JobQueue <- job
	}
	wg.Done()
}