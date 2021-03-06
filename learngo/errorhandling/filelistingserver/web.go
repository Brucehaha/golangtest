package main

import (
	"log"
	"net/http"
	"os"

	"brucego.com/learngo/errorhandling/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errorWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}

		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error occured "+"handling request: %s", err.Error())
			if userError, ok := err.(userError); ok {
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code),
				code)

		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {

	http.HandleFunc("/", errorWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
