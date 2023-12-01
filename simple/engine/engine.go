package engine

import (
	"shyCrawler/simple/fetch"
)

// SimpleEngine是一个简单的爬虫引擎，实现了对请求队列的处理和解析函数的调用

type SimpleEngine struct{}

// Run 启动引擎，并传入初始请求
func (s *SimpleEngine) Run(seeds ...Request) {
	// 初始化请求队列
	requests := make([]Request, 0)
	requests = append(requests, seeds...)

	// 不断从队列中取出请求并处理
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]

		// 处理请求，并获取解析结果
		parseResult, err := work(req)
		if err != nil {
			continue
		}

		// 将解析结果中的新请求加入队列
		requests = append(requests, parseResult.Requests...)
	}
}

// work 函数负责获取网页内容，并通过解析函数解析响应内容
func work(req Request) (ParseResult, error) {
	// 获取网页内容
	bytes, err := fetch.Fetch(req.Url)
	if err != nil {
		return ParseResult{}, err
	}

	// 解析网页内容，并返回解析结果
	return req.ParserFunc(bytes), nil
}
