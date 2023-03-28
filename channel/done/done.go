package main

//让main在 协程结束后才退出
import (
	"fmt"
	"sync"
)

const channelsSize = 10

// 直接打印出main goroutine和 doWork1 goroutine的通信
func chanDemo1() {
	var workers [channelsSize]worker1
	//create worker1
	for i := 0; i < len(workers); i++ {
		workers[i] = createWorker1(i)
	}
	//send msg
	for i, w := range workers {
		w.in <- fmt.Sprintf("《msg:%d》", i)
	}
	for _, w := range workers {
		<-w.done
	}

	for i, w := range workers {
		w.in <- fmt.Sprintf("《msg:%d》", i)
	}
	for _, w := range workers {
		<-w.done
	}
}

type worker1 struct {
	in   chan string
	done chan bool
}

func createWorker1(id int) worker1 {
	w := worker1{
		in:   make(chan string),
		done: make(chan bool),
	}
	go doWork1(id, w.in, w.done)
	return w
}

func doWork1(id int, c chan string, done chan bool) {
	for {
		fmt.Printf("Worker %d received %s\n", id, <-c)
		//发送给另一个channel进行goroutine通信
		//go func() {
		done <- true
		//}()
	}
}

// 使用标准库中的WaitGroup,类似java中的现成栅栏CyclicBarrier
func chanDemo2() {
	var workers [channelsSize]worker2
	var wg sync.WaitGroup
	//添加任务数量
	wg.Add(channelsSize * 2)
	//create worker1
	for i := 0; i < len(workers); i++ {
		workers[i] = createWorker2(i, &wg)
	}
	//send msg
	for i, w := range workers {
		w.in <- fmt.Sprintf("《msg:%d》", i)
	}

	for i, w := range workers {
		w.in <- fmt.Sprintf("《msg:%d》", i)
	}
	wg.Wait()
}

type worker2 struct {
	in   chan string
	done func()
}

func createWorker2(id int, wg *sync.WaitGroup) worker2 {
	w := worker2{
		in: make(chan string),
		done: func() {
			wg.Done()
		},
	}
	go doWork2(id, w)
	return w
}

func doWork2(id int, w worker2) {
	for {
		fmt.Printf("Worker %d received %s\n", id, <-w.in)
		//发送给另一个channel进行goroutine通信
		w.done()
	}
}

func main() {
	//chanDemo1()
	chanDemo2()
}
