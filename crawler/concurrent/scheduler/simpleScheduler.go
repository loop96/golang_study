package scheduler

import "golang_study/crawler/concurrent/engines"

type SimpleScheduler struct {
	workChannel chan engines.Request
}

// fixme:Why here use pointer receiver
func (s *SimpleScheduler) ConfigurationInChan(in chan engines.Request) {
	s.workChannel = in
}

func (s *SimpleScheduler) Submit(request engines.Request) {
	//must submit request to get result from out channel,so add goroutine here
	go func() { s.workChannel <- request }()
}
