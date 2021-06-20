package main

import (
	"net/http"
	"os"

	"brucego.com/learngo/errorhandling/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errorWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsNotExist(err):
				code = http.StatusForbidden
				//default:
				//	code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code),
				code)

		}
	}
}

func main() {

	http.HandleFunc("/list/", errorWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
