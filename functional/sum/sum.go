package main

import "fmt"

func sum() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	f := sum()
	for i := 0; i <= 10; i++ {
		fmt.Println(f(i))
	}
}
