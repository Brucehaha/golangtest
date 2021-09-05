package parser

import (
	"regexp"

	"brucego.com/learngo/crawler/engine"
)

const profileRe = `<a href="http://(album.zhenai.com/u/\d+)" [^>]*><img src=[^>]* alt="([^<]*)"></a>`

func ParseCity(content []byte) engine.ParseResult {
	re := regexp.MustCompile(profileRe)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: "https://" + string(m[1]),
			ParserFunc: func(b []byte) engine.ParseResult {
				return ParseProfile(b, string(m[2]))
			},
		})
	}

	return result

}
