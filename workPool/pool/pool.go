package pool

import (
	pJob "practice/workPool/job"
	"log"
)

type Work struct {
	Id int
	Job string
}

type Worker struct {
	Id int
	WorkerChannel chan chan Work
	Channel chan Work
	End chan bool
}

func (w *Worker) Start() {
	go func() {
		for{
			w.WorkerChannel <- w.Channel
			select {
			case job := <-w.Channel:
				pJob.DoWork(job.Job,w.Id)
			case <-w.End:
				return
			}
		}
	}()
}

// todo
func (w *Worker) Stop()  {
	log.Printf("work [%d] is stopping\n", w.Id)
	w.End <- true
}
