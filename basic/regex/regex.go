package main

import (
	"fmt"
	"regexp"
)

//todo
//正则表达式的使用
//FindAll
//FindAllSubmatch

const text = `<a href="http://www.zhenai.com/zhenghun/zhangjiakou" data-v-602e7f5e>张家口</a>`

func main() {
	compile := regexp.MustCompile(`<a href="http://www\.zhenai\.com/zhenghun/[a-zA-Z]+" [^>]*>[^<]+</a>`)
	match := compile.FindAllString(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
