package scheduler

import "golang_study/crawler/concurrent/engines"

type SimpleScheduler struct {
	WorkerChannel chan engines.Request
}

func (s *SimpleScheduler) GetWorkerChannel() chan engines.Request {
	return s.WorkerChannel
}

func (s *SimpleScheduler) Submit(request engines.Request) {
	//must submit request to get result from out channel,so add goroutine here
	go func() { s.WorkerChannel <- request }()
}

func (s *SimpleScheduler) WorkerReady(workerChannel chan engines.Request) {
}

func (s *SimpleScheduler) Run() {
	s.WorkerChannel = make(chan engines.Request)
}
