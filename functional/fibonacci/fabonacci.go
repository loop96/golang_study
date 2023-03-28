package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 定义一个函数来实现接口
type intGen func() int

// 实现Reader
func (i intGen) Read(p []byte) (n int, err error) {
	next := i()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%v\n", next)
	return strings.NewReader(s).Read(p)
}

func main() {
	f := fibonacci()
	printFileContents(f)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
