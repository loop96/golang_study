package main

import "fmt"

// GetMaxNonRepetitiveStr 最长不重复字符串
func GetMaxNonRepetitiveStr(str string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(str) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(GetMaxNonRepetitiveStr("as"))
	fmt.Println(GetMaxNonRepetitiveStr("cmvnczaad"))
	fmt.Println(GetMaxNonRepetitiveStr("abcasdabcffaaa"))
	fmt.Println(GetMaxNonRepetitiveStr("dasdasdff"))
}
