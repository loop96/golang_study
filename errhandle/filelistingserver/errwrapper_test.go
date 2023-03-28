package main

import (
	"errors"
	"fmt"
	"golang_study/errhandle/filelistingserver/filelisting"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func testUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func testNotFoundError(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func testNoPermissionError(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func testUnknownErr(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("idk")
}

func testNoErr(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintf(writer, "no err")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{filelisting.ErrHandler, 400, "path must start with /list/"},
	{filelisting.PanicHandler, 500, "Internal Server Error"},
	{testUserError, 400, "user error"},
	{testNotFoundError, 404, "Not Found"},
	{testNoPermissionError, 403, "Forbidden"},
	{testUnknownErr, 500, "Internal Server Error"},
	{testNoErr, 200, "no err"},
}

// 通过函数接口，使用假的request和response进行测试
func TestErrWrapper(t *testing.T) {
	for _, test := range tests {
		f := errWrapper(test.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
		f(response, request)

		VerifyResponse(t, response.Result(), test)
	}
}

// 真正起一个sever进行测试
func TestErrWrapperInServer(t *testing.T) {
	for _, test := range tests {
		f := errWrapper(test.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)

		VerifyResponse(t, response, test)
	}
}

func VerifyResponse(t *testing.T, response *http.Response, test struct {
	h       appHandler
	code    int
	message string
}) {
	bodyByte, _ := io.ReadAll(response.Body)
	body := strings.Trim(string(bodyByte), "\n")
	if response.StatusCode != test.code || body != test.message {
		t.Errorf("expect (%d,%s); got (%d,%s)",
			test.code, test.message, response.StatusCode, body)
	}
}
