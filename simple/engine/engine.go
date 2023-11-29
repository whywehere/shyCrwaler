package engine

import (
	"shyCrawler/simple/fetch"
)

type SimpleEngine struct{}

// Run runs the engine
func (s *SimpleEngine) Run(seeds ...Request) {
	requests := make([]Request, 0)
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]

		parseResult, err := work(req)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
	}
}

// work 获取网页信息
func work(req Request) (ParseResult, error) {
	bytes, err := fetch.Fetch(req.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return req.ParserFunc(bytes), nil
}
