package main

import (
	"shyCrawler/simple/engine"
	"shyCrawler/simple/parser"
)

const URL = "https://book.douban.com/"

func main() {
	e := engine.SimpleEngine{}
	e.Run(engine.Request{
		Url:        URL,
		ParserFunc: parser.ParseBookList,
	})
}
