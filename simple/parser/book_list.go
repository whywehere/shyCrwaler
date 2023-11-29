package parser

import (
	"regexp"
	"shyCrawler/simple/engine"
)

const bookListRe = `<a href="([^"]+)" title="([^"]+)"`

func ParseBookList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(bookListRe)
	subMatch := re.FindAllSubmatch(contents, -1) // -1 表示匹配所有符合的项
	result := engine.ParseResult{}
	for _, m := range subMatch {
		url := string(m[1])
		name := string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(content []byte) engine.ParseResult {
				return ParseBookDetail(content, name)
			},
		})
	}
	return result
}
