package engines

import (
	"golang_study/crawler/standalone/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemSaver chan interface{}
}

type Scheduler interface {
	Submit(request Request)
	// GetWorkerChannel To get current used worker channel from engine
	GetWorkerChannel() chan Request
	WorkerNotifier
	Run()
}

type WorkerNotifier interface {
	WorkerReady(workerChannel chan Request)
}

func (e ConcurrentEngine) Run(seed ...Request) {
	//create channel to sync result
	out := make(chan ParserResult)
	e.Scheduler.Run()
	in := e.Scheduler.GetWorkerChannel()
	//must create worker to work
	for i := 0; i < e.WorkCount; i++ {
		// use the channel to get data or put in data
		e.CreateWorker(in, out, e.Scheduler)
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
			go func(item interface{}) { e.ItemSaver <- item }(item)
		}
		//put in request,here dependence result from out channel,so it could be stuck if not use goroutine
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func (e ConcurrentEngine) Worker(r Request) (ParserResult, error) {
	log.Printf("fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch err,url:%s:%v", r.Url, err)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}

func (e ConcurrentEngine) CreateWorker(in chan Request, out chan ParserResult, w WorkerNotifier) {
	go func() {
		for {
			w.WorkerReady(in)
			request := <-in
			result, err := e.Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
