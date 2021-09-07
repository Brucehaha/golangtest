package engine

import (
	"log"

	"brucego.com/learngo/crawler/fetcher"
)

var headers = map[string]string{
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"cookie":                    "sid=c215a129-1cca-458c-a1fe-3b9bf64f7855; ec=50lv3h6O-1630743699780-d3e1a55a9ed98673742148; FSSBBIl1UgzbN7NO=5Uh_xj4kFKbkktmPkxqsz51jvjXid91WlHQcJEMgoM7QwfEtuVYnRbi9FxfaRw_wwfoB3bqLXNgyBvvvVWxvQtA; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1630743804,1630816283,1630925331,1630933733; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1630933746; _exid=qtfyfypSv6fKTOb6pinUxT7AFn%2BOlId2KRkMm5bHyRDbZtgw0Ui13kHI%2F1eJVp9a9cLqUE1fe%2BZJxTplriE4jQ%3D%3D; _efmdata=19qjfAjF9xJWgvF7jOi9XOqFk97mxsF5ucC7N%2B3MsS1eRMCH2yQCAokmYdDLJfGvQFxZwH%2FNX4tSNFe8qKIJTckXTaN5zRcDpaOiAqjDCro%3D; FSSBBIl1UgzbN7NP=53s3sqblClB0qqqm4XuNr1axuC2zprLb8XnrS0x5XyE7r1D3dAlU88hSUz8hXCJTapSyjT7BhCDcAn1ZobniZ4PtpDKBv1v381xmDq._3d3ziiCOfkjwHi3iQBgDJ331yztQ0_RvFn.4VX.jmdNGyLboARTgGyQ6D_aTnNZFjiMJZ60lNjvh6vesLpDRwaCsd20x2IkRgXIvG8WqQU9AFDy7Gk1pvYDPziyBr8mG7Qy6hcS3bd49fv7oZbj7A.9fQG4vbRD7S1pGJ3xKvR3JQtZ",
	"accept-language":           "en-GB,en;q=0.9,en-US;q=0.8",
	"cache-control":             "max-age=0",
	"referer":                   " https://album.zhenai.com/",
	"sec-ch-ua":                 `Microsoft Edge";v="93", " Not;A Brand";v="99", "Chromium";v="93`,
	"sec-ch-ua-mobile":          "?0",
	"sec-ch-ua-platform":        "Windows",
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "same-site",
	"sec-fetch-user":            "?1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36 Edg/93.0.961.38",
}

func Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url, headers)
		if err != nil {
			log.Printf("error with url %s: %v", r.Url, err)
		}
		results := r.ParserFunc(body)
		requests = append(requests, results.Requests...)
		for _, i := range results.Items {
			log.Printf("Got result %v", i)
		}

	}
}
