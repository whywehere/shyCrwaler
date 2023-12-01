package engine

// Request 是一个表示爬虫请求的结构体，包含了请求的URL和解析函数
type Request struct {
	Url        string                   // 请求的URL
	ParserFunc func([]byte) ParseResult // 解析函数，用于解析响应内容
}

// ParseResult 表示解析结果的结构体，包含了新的请求和解析出的数据
type ParseResult struct {
	Requests []Request     // 新的请求
	Items    []interface{} // 解析出的数据
}

// NilParser 是一个空解析函数，用于初始化Request的ParserFunc字段
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
