package fetcher

//noinspection GoUnresolvedReference
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"golang.org/x/net/html/charset"
)

func Fetch(url string, headers map[string]string) ([]byte, error) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %v ", resp.StatusCode)
	}

	bufioReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bufioReader)
	newBody := transform.NewReader(bufioReader, e.NewEncoder())
	return ioutil.ReadAll(newBody)

}
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}
