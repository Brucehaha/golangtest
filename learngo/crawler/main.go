package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"

	"golang.org/x/text/encoding"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Print("error: status code", resp.StatusCode)
		return
	}
	bufioReader := bufio.NewReader(resp.Body)
	_, name := determineEncoding(bufioReader)
	//newBody := transform.NewReader(resp.Body, e.NewEncoder())
	fmt.Println(name)
	all, err := ioutil.ReadAll(bufioReader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)

}
func determineEncoding(r *bufio.Reader) (e encoding.Encoding, name string) {
	bytes, err := r.Peek(1024)
	if err != nil {

		panic(err)
	}
	e, names, _ := charset.DetermineEncoding(bytes, "")
	return e, names
}
