package main

import (
	"golang_study/crawler/standalone/engines"
	"golang_study/crawler/standalone/scrapeCenter/parser"
)

// 并发爬虫
// 1.在单机的基础上，将最耗时的fetch与parse和部分engine构成worker
// 2.加入并发调度器scheduler来分配worker任务
func main() {
	//engines.Run([]engines.Request{
	//	{
	//		Url:        "https://www.zhenai.com/zhenghun",
	//		ParserFunc: parser.ParseCityList,
	//	},
	//})

	engines.Run([]engines.Request{
		{
			Url:        parser.BaseUrl,
			ParserFunc: parser.ParsePageList,
		},
	})
}
