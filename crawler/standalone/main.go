package main

import (
	"golang_study/crawler/standalone/engines"
	"golang_study/crawler/standalone/scrapeCenter/parser"
)

// 单机爬虫
// 1.使用正则表达式爬取数据
// 2.使用FindAllSubmatch获取城市跳转地址和城市名称
// 3.页面获取 fetch
// 4.解析器 Parse**
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
