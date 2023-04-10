package main

import (
	"golang_study/crawler/concurrent/engines"
	"golang_study/crawler/concurrent/persist"
	"golang_study/crawler/concurrent/scheduler"
	"golang_study/crawler/concurrent/scrapeCenter/parser"
)

// 并发爬虫
// 1.在单机的基础上，将最耗时的fetch与parse和部分engine构成worker
// 2.加入并发调度器scheduler来分配worker任务
// 3.加入队列调度器
func main() {
	e := engines.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 100,
		ItemSaver: persist.GetItemSaver(),
	}
	e.Run(engines.Request{
		Url:        parser.BaseUrl,
		ParserFunc: parser.ParsePageList,
	})
}
