package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func printFile(filePath string) {

	//file, err := os.Open(filePath)
	//if err != nil {
	//	panic(err)
	//}

	s := `123123
adsdasdasdasdff
ddmcmcmnncnc
""asd91230-0=-='''da'sd'as'd''ads""`
	reader := strings.NewReader(s)
	printFileContents(reader)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	printFile("/Users/panyi/Documents/GolandProjects/golang_study/io/123.txt")
}
