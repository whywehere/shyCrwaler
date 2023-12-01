package parser

import (
	"regexp"
	"shyCrawler/simple/engine"
)

const bookListRe = `<a href="([^"]+)" title="([^"]+)"` // 正则表达式，用于匹配书籍列表中的URL和书名

// ParseBookList 用于解析书籍列表页面的内容，提取书籍的URL和书名，并生成对应的解析结果
func ParseBookList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(bookListRe)         // 编译正则表达式
	subMatch := re.FindAllSubmatch(contents, -1) // 使用正则表达式进行匹配，-1 表示匹配所有符合的项
	result := engine.ParseResult{}               // 创建一个解析结果实例

	for _, m := range subMatch {
		url := string(m[1])  // 提取书籍详情页面的URL
		name := string(m[2]) // 提取书籍名称

		result.Items = append(result.Items, name) // 将书籍名称放入解析结果的 Items 字段中

		// 创建一个新的请求，请求URL为书籍详情页面的URL，解析函数为 ParseBookDetail
		request := engine.Request{
			Url: url,
			ParserFunc: func(content []byte) engine.ParseResult {
				return ParseBookDetail(content, name) // 调用 ParseBookDetail 解析书籍详情页面
			},
		}
		result.Requests = append(result.Requests, request) // 将请求放入解析结果的 Requests 字段中
	}

	return result // 返回解析结果
}
