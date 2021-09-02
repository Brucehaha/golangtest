package fetcher

//noinspection GoUnresolvedReference
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"

	"golang.org/x/net/html/charset"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %v ", resp.StatusCode)
	}
	bufioReader := bufio.NewReader(resp.Body)
	e, _ := determineEncoding(bufioReader)
	newBody := transform.NewReader(resp.Body, e.NewEncoder())
	return ioutil.ReadAll(newBody)

}
func determineEncoding(r *bufio.Reader) (e encoding.Encoding, name string) {
	bytes, err := r.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, names, _ := charset.DetermineEncoding(bytes, "")
	return e, names

}
