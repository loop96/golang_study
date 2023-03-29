package engines

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(request Request)
	ConfigurationInChan(in chan Request)
}

func (e ConcurrentEngine) Run(seed ...Request) {

	//create channel to sync request&result
	in := make(chan Request)
	e.Scheduler.ConfigurationInChan(in)
	out := make(chan ParserResult)

	//must create worker to work
	for i := 0; i < e.WorkCount; i++ {
		// use the channel to get data or put in data
		e.CreateWorker(in, out)
	}

	//init request
	for _, r := range seed {
		e.Scheduler.Submit(r)
	}

	//real run body
	for {
		//get data
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item %v", item)
		}
		//put in request,here dependence result from out channel,so it could be stuck if not use goroutine
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func (e ConcurrentEngine) CreateWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
