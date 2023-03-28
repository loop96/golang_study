package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Euler 欧拉函数，输出(0+1.2246467991473515e-16i)是因为e和π都是float64，跟python一样的
// %.3f 输出即为0
func Euler() complex128 {
	return cmplx.Pow(math.E, 1i*math.Pi) + 1
}

// Enums 常量类型，iota说明从0开始，后面的常量自增1
func Enums() {
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
	const (
		B = 1 << (10 * iota)
		KB
		MB
		GB
		TB
	)
	fmt.Println(B, KB, MB, GB, TB)
}

func main() {
	fmt.Println(Euler())
	Enums()
}
