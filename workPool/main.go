package main

import (
	"log"
	"practice/workPool/pool"
	pJob "practice/workPool/job"
)

const WORKER_COUNT = 5
const JOB_COUNT = 100

// 入口函数
func main() {
	log.Println("starting application...")
	collector := pool.StartDispatcher(WORKER_COUNT)
	for i,job := range pJob.CreateJobs(JOB_COUNT){
		collector.Work <- pool.Work{Job:job,Id:i}
	}
}
