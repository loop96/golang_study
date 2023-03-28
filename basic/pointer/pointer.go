package main

import "fmt"

// swap_val 值传递，错误方式
func swap_val(a, b int) {
	a, b = b, a
}

// swap_pointer_val 使用指针的地址值传递，正确但不规范
func swap_pointer_val(a, b *int) {
	*a, *b = *b, *a
}

// swap_normative 值传递，返回结果
func swap_normative(a, b int) (int, int) {
	return b, a
}

func PrintSwap() {
	a, b := 1, 2
	c, d := 1, 2
	swap_val(a, b)
	fmt.Printf("swap_val a=%d,b=%d\n", a, b)
	swap_pointer_val(&a, &b)
	fmt.Printf("swap_pointer_val a=%d,b=%d\n", a, b)
	c, d = swap_normative(c, d)
	fmt.Printf("swap_normative a=%d,b=%d\n", c, d)
}

func main() {
	PrintSwap()
}
