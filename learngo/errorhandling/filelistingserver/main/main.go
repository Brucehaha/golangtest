package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/list", func(writer http.ResponseWriter, request *http.Request) {

	})
}
