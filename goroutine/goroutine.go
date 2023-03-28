package main

import (
	"fmt"
	"time"
)

// goroutine 非抢占，如果协程内没有人让出时间片其他协程不会抢用线程资源
// fmt.printf是IO操作，会有一段进行IO从而让出时间片
func example1() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("fuck_world = %v\n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}

// go 1.14后优化了这个问题，不会出现死循环
func example2() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

// 使用闭包的方式写的话，闭包函数中的i关联着for里面的i，在最后一个循环中为10，a[10]越界
// 使用go run -race .
// WARNING: DATA RACE
// 读
// Read at 0x00c00013e018 by goroutine 7:
//
//	main.example3.func1()
//	    /Users/panyi/Documents/GolandProjects/golang_study/goroutine/goroutine.go:43 +0x60
//
// 上次一的写
// Previous write at 0x00c00013e018 by main goroutine:
//
//	main.example3()
//	    /Users/panyi/Documents/GolandProjects/golang_study/goroutine/goroutine.go:39 +0x88
//	main.main()
//	    /Users/panyi/Documents/GolandProjects/golang_study/goroutine/goroutine.go:52 +0x20
//
// 可得知for循环的i在a[i]前读写发生冲突
func example3() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func() {
			for {
				//由 'go' 语句中的 'func' 文字捕获的循环变量可能有意外值
				a[i]++
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func main() {
	//example1()
	example2()
	//example3()
}
