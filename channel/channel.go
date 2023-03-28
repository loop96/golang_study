package main

import (
	"fmt"
	"time"
)

// all goroutines are asleep - deadlock!
func channelDemoError() {
	channel := make(chan int)
	channel <- 1
	channel <- 2
	fmt.Println(<-channel)
}

// 使用make声明一个带有buffer的channel
func channelDemo1() {
	//buffer
	channel := make(chan int, 10)
	channel <- 1
	channel <- 2
	fmt.Println(<-channel)
}

// 再开一个goroutine死循环去拿channel的数据
func channelDemo2() {
	channel := make(chan int)
	go func() {
		for {
			fmt.Println(<-channel)
		}
	}()
	channel <- 1
	channel <- 2
	//防止方法退出后程序也退出导致上面goroutine没打印出来
	time.Sleep(time.Millisecond)
}

const channelsSize = 10

func channelDemo3() {
	var channels [channelsSize]chan string
	//create worker
	for i := 0; i < channelsSize; i++ {
		channels[i] = createWorker(i)
	}
	//send msg
	for i := 0; i < channelsSize; i++ {
		channels[i] <- fmt.Sprintf("msg:%d", i)
	}
	for i := 0; i < channelsSize; i++ {
		channels[i] <- fmt.Sprintf("msg:%d", i)
	}
	//防止方法退出后程序也退出导致上面goroutine没打印出来
	time.Sleep(time.Millisecond)
}

func createWorker(id int) chan string {
	c := make(chan string)
	go func() {
		// chan range
		for v := range c {
			fmt.Printf("Worker %d received %s\n", id, v)
		}
	}()
	return c
}

// chan的固定申明
func channelDemo4() {
	//只能收
	//channelReceiver := make(chan<- int)
	//只能发
	//channelSender := make(<-chan int)
	//channelSender <- 1 err!
	//channelReceiver <- 1
	//fmt.Println(<-channelReceiver) err!
	//fmt.Println(<-channelSender)
}

// 如果channel被关闭了 取出来的都会是channel里面类型的零值
func channelDemo5() {
	channel := make(chan int)
	go func() {
		for {
			v, ok := <-channel
			if !ok {
				break
			}
			fmt.Printf("received %d\n", v)
		}
	}()
	channel <- 1
	close(channel)
	time.Sleep(time.Millisecond)
}

func main() {
	//channelDemoError() err!
	//channelDemo1()
	//channelDemo2()
	//channelDemo3()
	//channelDemo4()
	//channelDemo5()
}
