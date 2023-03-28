package fetcher

import (
	"bufio"
	"errors"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Fetch 根据地址去获取body
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Printf("Request err: %v", err)
		return nil, errors.New(strings.Join([]string{"Error:status Code="}, strconv.Itoa(resp.StatusCode)))
	}
	//determineEncoding自动去解码
	bodyReader := bufio.NewReader(resp.Body)
	e, err := determineEncoding(bodyReader)
	if err != nil {
		return nil, err
	}
	reader := transform.NewReader(bodyReader, e.NewDecoder())
	all, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	return all, nil
}

func determineEncoding(r *bufio.Reader) (encoding.Encoding, error) {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetch err: %v", err)
		return unicode.UTF8, err
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e, nil
}
