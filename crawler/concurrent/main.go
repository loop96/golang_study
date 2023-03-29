package main

import (
	"golang_study/crawler/concurrent/engines"
	"golang_study/crawler/concurrent/scheduler"
	"golang_study/crawler/concurrent/scrapeCenter/parser"
)

// 并发爬虫
// 1.在单机的基础上，将最耗时的fetch与parse和部分engine构成worker
// 2.加入并发调度器scheduler来分配worker任务
// 2.1.简单的调度器只使用一个channel用来接收数据,一个channel用来发送数据
// 2.2.并发的调度器
// todo stuck in movieDetail
func main() {
	e := engines.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkCount: 10,
	}
	e.Run(engines.Request{
		Url:        parser.BaseUrl,
		ParserFunc: parser.ParsePageList,
	})
}
