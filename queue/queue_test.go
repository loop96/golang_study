package queue

import "fmt"

func ExampleQueue_Poll() {
	q := new(Queue)
	q.Push(1)
	q.Push("3str")
	fmt.Println(q.Poll())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Poll())
	fmt.Println(q.IsEmpty())
	// Output:
	//1
	//false
	//3str
	//true
}
