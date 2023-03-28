package main

import (
	"testing"
)

// 因为go的语法特殊，适合使用表格驱动测试
// 命令行 go test .
// 覆盖率命令行  go test -coverprofile=c.out
// go tool cover -html=c.out
func TestGetMaxNonRepetitiveStr(t *testing.T) {
	tests := []struct {
		a int
		b string
	}{
		{2, "as"},
		{6, "cmvnczaad"},
		{6, "abcasdabcffaaa"},
		{4, "dasdasdff"},
		{10, "阿是度IC哦UI正常年对的卡卡拉奇怕怕"},
	}
	for _, test := range tests {
		if actual := GetMaxNonRepetitiveStr(test.b); actual != test.a {
			t.Errorf("GetMaxNonRepetitiveStr(%s) got %d ,execpted %d\n", test.b, actual, test.a)
		}
	}

}

// 性能测试
// 命令行 go test -bench .
func BenchmarkGetMaxNonRepetitiveStr(b *testing.B) {
	s := "阿是度IC哦UI正常年对的卡卡拉奇怕怕"
	ans := 10
	//b.N系统自行定义
	for i := 0; i < b.N; i++ {
		if actual := GetMaxNonRepetitiveStr(s); actual != ans {
			b.Errorf("GetMaxNonRepetitiveStr(%s) got %d ,execpted %d\n", s, actual, ans)
		}
	}
}
