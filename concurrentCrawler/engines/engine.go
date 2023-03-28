package engines

import (
	"golang_study/crawler/standalone/fetcher"
	"log"
)

func Run(requests []Request) {
	//循环request进行调用fetcher
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetch err,url:%s:%v", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		//直接将parseResult里的request全放进去 类似addAll
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			//TODO
			log.Printf("Got item %v", item)
		}
	}

}
