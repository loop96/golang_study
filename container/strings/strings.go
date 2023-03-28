package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "WTF搞几把呢！?" //UTF-8
	fmt.Println(str, len(str))

	for i, ch := range []byte(str) {
		fmt.Print(i, ch)
	}
	fmt.Println()
	for i, r := range str {
		fmt.Print(i, r)
	}
	fmt.Println()
	count := utf8.RuneCountInString(str)
	fmt.Println(count)

	//正确处理国际化字符串
	for i, r := range []rune(str) {
		fmt.Printf("(%d = %c)", i, r)
	}
	fmt.Println()
	fmt.Println(strings.ContainsAny(str, "几啊啊实打实多撒打算"))
	fmt.Println(strings.Contains(str, "几把呢"))
	//r, _ := utf8.DecodeRune([]byte("几"))
	//fmt.Println(strings.ContainsRune(str, r))
}
