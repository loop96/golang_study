package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(prefix string) chan string {
	c := make(chan string)
	i := 1
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("%s%d", prefix, i)
			i++
		}
	}()
	return c
}

func createWorker(id int) chan string {
	c := make(chan string)
	go worker(c, id)
	return c
}

// 在此充当消费者
func worker(c chan string, id int) {
	// chan range
	for v := range c {
		//模拟消费者消费时间
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d received %v\n", id, v)
	}
}

func main() {
	//demo1()
	demo2()

}

// 使用信号和nil chan阻塞来模拟生产消费者模型，但是消费者如果消费过慢，c1/c2的v值会被覆盖
func demo1() {
	c1, c2 := generator("c1_"), generator("c2_")
	var worker = createWorker(0)
	v := ""
	hasValue := false
	for {
		//case中的谁先接收谁先执行，不加default则阻塞，加了default是非阻塞
		var activateWorker chan string //nil chan 在select是一直被阻塞的
		if hasValue {
			activateWorker = worker
		}
		select {
		case v = <-c1:
			hasValue = true
		case v = <-c2:
			hasValue = true
		case activateWorker <- v:
			//当c1和c2生产出一个值后，放入消费者，并把标识置为false
			hasValue = false
		}
	}
}

// 使用queue来做缓冲队列
func demo2() {
	c1, c2 := generator("c1_"), generator("c2_")
	var worker = createWorker(0)
	//声明一个队列，用来缓存demo1中c1/c2因为消费者消费过慢覆盖的数据
	var queue []string
	//声明一个定时chan，可以来操作生产消费循环
	after := time.After(10 * time.Second)
	//每秒向tick channel塞数据
	tick := time.Tick(time.Second)
	for {
		var activateWorker chan string //nil chan 在select是一直被阻塞的
		var activateValue string
		if len(queue) > 0 {
			//如果队列不为空，取出来的值塞入worker中供消费
			activateWorker = worker
			activateValue = queue[0]
		}
		select {
		case v := <-c1:
			//如果c1有数据进来，添加到队列
			queue = append(queue, v)
		case v := <-c2:
			//如果c2有数据进来，添加到队列
			queue = append(queue, v)
		case activateWorker <- activateValue:
			//交给worker消费后去掉queue的元素
			queue = queue[1:]
			fmt.Println(queue)
		case <-time.After(500 * time.Millisecond):
			//0.5s 内消费者没有消费完，生产者没生产，timeout
			fmt.Println("timeout")
		case <-tick:
			fmt.Println(queue)
		case <-after:
			fmt.Println("bye")
			return
		}
	}
}
