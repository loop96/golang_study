package scheduler

import "golang_study/crawler/concurrent/engines"

type QueueScheduler struct {
	RequestChannel chan engines.Request
	WorkerChannel  chan chan engines.Request
}

func (q *QueueScheduler) Submit(request engines.Request) {
	q.RequestChannel <- request
}

func (q *QueueScheduler) WorkerReady(workerChannel chan engines.Request) {
	q.WorkerChannel <- workerChannel

}

func (q *QueueScheduler) GetWorkerChannel() chan engines.Request {
	return make(chan engines.Request)
}

func (q *QueueScheduler) Run() {
	q.RequestChannel = make(chan engines.Request)
	q.WorkerChannel = make(chan chan engines.Request)
	go func() {
		var requestQ []engines.Request
		var workerQ []chan engines.Request

		for {
			var activateRequest engines.Request
			var activateWorker chan engines.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activateRequest = requestQ[0]
				activateWorker = workerQ[0]
			}
			select {
			case w := <-q.WorkerChannel:
				// send next request to w
				workerQ = append(workerQ, w)
			case r := <-q.RequestChannel:
				// send request to worker
				requestQ = append(requestQ, r)
			case activateWorker <- activateRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()

}
