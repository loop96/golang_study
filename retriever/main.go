package main

import (
	"fmt"
	"golang_study/retriever/mock"
	"golang_study/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "https://www.baidu.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(p Poster) {
	p.Post(url,
		map[string]string{
			"name":   "test",
			"course": "golang",
		})
}

// RetrieverPoster interface 的组合
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(rp RetrieverPoster) string {
	rp.Post(url, map[string]string{
		"Contents": "another fake",
	})
	return rp.Get(url)
}

func main() {
	var r Retriever
	mockRetriever := mock.Retriever{Contents: "this is fake"}
	r = &mockRetriever
	inspect(r)
	//mockRetriever := r.(mock.Retriever)
	//fmt.Println(mockRetriever.Contents)
	//fmt.Println(download(r))

	//指针接收者的接口实现
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Time:      time.Minute,
	}
	inspect(r)
	//mockRetriever = r.(mock.Retriever) //panic: interface conversion: main.Retriever is *real.Retriever, not mock.Retriever
	//realRetriever := r.(*real.Retriever)
	//fmt.Println(realRetriever.UserAgent)
	//fmt.Println(download(r))

	fmt.Println("==Type assertion")
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println("Contents:", mockRetriever.Contents)
	} else {
		fmt.Println("not a mock Retriever")
	}
	fmt.Println("==Try session")
	fmt.Println(session(&mockRetriever))

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" >%T %v\n", r, r)
	fmt.Print(" >Type switch")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
