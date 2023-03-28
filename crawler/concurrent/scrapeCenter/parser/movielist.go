package parser

import (
	"golang_study/crawler/standalone/engines"
	"regexp"
)

const BaseUrl = "https://ssr1.scrape.center"

var movieReg = regexp.MustCompile(`<a\s+[a-zA-Z0-9\-]+=""\s+href="(/detail/\d+)"\s+class="name"[^>]*>\s+<[^<]*h2\s+[a-zA-Z0-9\-]+=""\s+class="[a-zA-Z0-9\-]+">(.+)</h2>\s+</a>`)

// ParseMovieList 解析列表
func ParseMovieList(htmlByte []byte) engines.ParserResult {
	match := movieReg.FindAllSubmatch(htmlByte, -1)
	result := engines.ParserResult{}
	for _, m := range match {
		url := string(m[1])
		name := string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engines.Request{
			Url:        BaseUrl + url,
			ParserFunc: ParseMovieDetail,
		})
	}
	return result
}
