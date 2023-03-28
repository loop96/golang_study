package main

import (
	engines2 "golang_study/crawler/standalone/engines"
	parser2 "golang_study/crawler/standalone/scrapeCenter/parser"
)

// 单机爬虫
// 1.使用正则表达式爬取数据
// 2.使用FindAllSubmatch获取城市跳转地址和城市名称
// 3.页面获取 fetch
// 4.解析器 Parse
func main() {
	//engines.Run([]engines.Request{
	//	{
	//		Url:        "https://www.zhenai.com/zhenghun",
	//		ParserFunc: parser.ParseCityList,
	//	},
	//})

	engines2.Run([]engines2.Request{
		{
			Url:        parser2.BaseUrl,
			ParserFunc: parser2.ParsePageList,
		},
	})
}
