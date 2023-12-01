package main

import (
	"shyCrawler/simple/engine"
	"shyCrawler/simple/parser"
)

const URL = "https://book.douban.com/"

func main() {
	// 创建一个简单引擎对象
	e := engine.SimpleEngine{}

	// 启动引擎，传入初始请求
	e.Run(engine.Request{
		Url:        URL,                  // 初始请求的URL
		ParserFunc: parser.ParseBookList, // 解析函数，用于解析响应内容
	})
}
