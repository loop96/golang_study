package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func Apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling funciton %s with args %d,%d \n", opName, a, b)
	return op(a, b)
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(Apply(Pow, 3, 3))
}
