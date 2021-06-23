package filelisting

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}
func (e userError) Message() string {
	return string(e)
}
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path
	if strings.Index(path, prefix) != 0 {
		return userError("paths must start with " + prefix)
	}
	fmt.Println(len(path), reflect.TypeOf(path), path)

	path = path[len(prefix):]

	log.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
