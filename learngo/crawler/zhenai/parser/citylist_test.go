package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	content, err := ioutil.ReadFile("citylist.html")
	if err != nil {
		panic(err)
	}
	results := ParseCityList(content)
	cities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	urls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/akesu",
	}

	const resultSize = 470
	if len(results.Requests) != resultSize {
		t.Errorf("city list url size should be %v, not  %d", len(results.Requests), resultSize)
	}
	for i, v := range urls {
		if results.Requests[i].Url != v {
			t.Errorf("url should be %s", results.Requests[i].Url)
		}

	}

	if len(results.Items) != resultSize {
		t.Errorf("city list city name size should be %v, not  %d", len(results.Items), resultSize)
	}
	for i, v := range cities {
		if results.Items[i] != v {
			t.Errorf("url should be %s", results.Items[i])
		}

	}

}
