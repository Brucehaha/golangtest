package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	content, err := ioutil.ReadFile("test_fixture/cities.html")
	if err != nil {
		panic(err)
	}
	items := []string{
		"请你做我人间欢喜", "赫语", "Ainee",
	}
	urls := []string{
		"http://album.zhenai.com/u/1690425369", "http://album.zhenai.com/u/1394334012", "http://album.zhenai.com/u/1518285094",
	}

	results := ParseCity(content)
	for i, v := range items {
		if results.Items[i] != v {
			t.Error(results.Items[i])
		}
	}
	for i, v := range urls {
		url := results.Requests[i].Url
		if url != v {
			t.Error(url)
		}
	}

}
