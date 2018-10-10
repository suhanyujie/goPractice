package pool

import "log"

type Collector struct {
	Work chan Work
	End  chan bool
}

var WorkerChannel = make(chan chan Work)

func StartDispatcher(workerCount int) Collector {
	var i int
	var workers []Worker
	input := make(chan Work)
	end := make(chan bool)
	collector := Collector{Work: input, End: end}
	for i < workerCount {
		i++
		log.Println("start working:", i)
		worker := Worker{
			Id:            i,
			Channel:       make(chan Work),
			WorkerChannel: WorkerChannel,
			End:           make(chan bool),
		}
		worker.Start()
		workers = append(workers, worker)
	}

	go func() {
		for {
			select {
			case <-end:
				for _, w := range workers {
					w.Stop()
				}
				return
			case work := <-input:
				worker := <-WorkerChannel
				worker <- work
			}

		}
	}()
	return collector
}
