package engine

import (
	"log"

	"brucego.com/learngo/crawler/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)
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
