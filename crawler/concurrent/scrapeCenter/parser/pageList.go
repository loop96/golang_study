package parser

import (
	"golang_study/crawler/standalone/engines"
	"log"
	"regexp"
	"strconv"
)

var movieListSizeReg = regexp.MustCompile(`<a href="(\/page\/\d+)">(\d+)<\/a>`)
var movieListSumReg = regexp.MustCompile(`<span class="el-pagination__total">共 (\d+) 条</span>`)

// ParsePageList 页数解析器
func ParsePageList(htmlByte []byte) engines.ParserResult {
	//获取当前页码和总页码
	sumMatch := movieListSumReg.FindSubmatch(htmlByte)
	totalNum, err := strconv.Atoi(string(sumMatch[1]))
	if err != nil {
		panic(err)
	}
	match := movieListSizeReg.FindAllSubmatch(htmlByte, -1)
	if err != nil {
		panic(err)
	}
	log.Printf("get total %d data, size = %d \n ", totalNum, len(match))
	result := engines.ParserResult{}
	//构造待爬的每页数据
	for _, m := range match {
		pageParam := string(m[1])
		result.Items = append(result.Items, pageParam)
		result.Requests = append(result.Requests, engines.Request{
			Url:        BaseUrl + pageParam,
			ParserFunc: ParseMovieList,
		})
	}
	return result
}
