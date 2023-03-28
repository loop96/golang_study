package main

import (
	"fmt"
	"golang_study/queue"
)

func main() {
	q := new(queue.Queue)
	q.Push(1)
	q.Push("3str")
	fmt.Println(q.Poll())
	fmt.Println(q.Peek())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Poll())
	fmt.Println(q.IsEmpty())
}
