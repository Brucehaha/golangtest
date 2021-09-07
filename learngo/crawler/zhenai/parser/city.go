package parser

import (
	"fmt"
	"regexp"

	"brucego.com/learngo/crawler/engine"
)

const profileRe = `<a href="http://album.zhenai.com/u/(\d+)" [^>]*><img src=[^>]* alt="([^<]*)"></a>`

func ParseCity(content []byte) engine.ParseResult {
	re := regexp.MustCompile(profileRe)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range matches {

		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: fmt.Sprintf(`https://album.zhenai.com/api/profile/getObjectProfile.do?objectID=%s&_=1630933757923&ua=h5%%2F1.0.0%%2F1%%2F0%%2F0%%2F0%%2F0%%2F0%%2F%%2F0%%2F0%%2Fe44e7713-ccb9-4449-8ef3-2adbb09f778c%%2F0%%2F0%%2F1023356863`, string(m[1])),
			ParserFunc: func(b []byte) engine.ParseResult {
				return ParseProfile(b, string(m[2]))
			},
		})
	}

	return result

}
