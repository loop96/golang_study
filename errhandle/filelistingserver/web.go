package main

import (
	"golang_study/errhandle/filelistingserver/filelisting"
	"log"
	"net/http"
	//使用{basePath}/debug/pprof/ 访问
	_ "net/http/pprof"
	"os"
)

// 真正的业务处理
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 错误包装
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.RequestURI() == "/favicon.ico" {
			return
		}

		//panic
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				code := http.StatusInternalServerError
				http.Error(writer, http.StatusText(code), code)
			}
		}()

		err := handler(writer, request)
		if nil != err {
			log.Printf("Error handler request: %s", err.Error())

			//user err
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			//system err
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.ErrHandler))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}

}
