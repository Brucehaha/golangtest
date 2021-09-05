package parser

import (
	"regexp"

	"brucego.com/learngo/crawler/engine"
)

const cityRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^>]+)</a>`

func ParseCityList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	limit := 10
	for _, m := range matches {
		if limit == 0 {
			return result
		}
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		limit--
	}

	return result

}
