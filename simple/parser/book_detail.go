package parser

import (
	"fmt"
	"log"
	"regexp"
	"shyCrawler/simple/engine"
	"shyCrawler/simple/model"
	"strconv"
)

var (
	authorRe    = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)                  // 作者的正则表达式
	publisherRe = regexp.MustCompile(`<span class="pl">出版社:</span>[\d\D]*?<a.*?>([^<]+)</a>`)                  // 出版社的正则表达式
	pagesRe     = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)                               // 页数的正则表达式
	priceRe     = regexp.MustCompile(`<span class="pl">定价:</span> ([^<]+)<br/>`)                               // 定价的正则表达式
	scoreRe     = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+) </strong>`) // 评分的正则表达式
	introRe     = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)                        // 简介的正则表达式
)

// ParseBookDetail 用于解析书籍详情页面的内容，提取相关信息并生成对应的解析结果
func ParseBookDetail(contents []byte, name string) engine.ParseResult {
	book := model.Book{} // 创建一个 Book 实例
	book.Name = name     // 设置书名

	// 通过 extraContent 函数提取作者、简介、页数、定价、出版社和评分等信息，并设置到 Book 实例中
	book.Author = extraContent(contents, authorRe)
	book.Intro = extraContent(contents, introRe)
	if pages, err := strconv.Atoi(extraContent(contents, pagesRe)); err == nil {
		book.Pages = pages
	} else {
		log.Print(err.Error())
	}
	book.Price = extraContent(contents, priceRe)
	book.Publisher = extraContent(contents, publisherRe)
	if score, err := strconv.ParseFloat(extraContent(contents, scoreRe), 64); err == nil {
		book.Score = score
	} else {
		log.Print(err.Error())
	}

	result := engine.ParseResult{
		Requests: nil,
		Items:    []interface{}{book}, // 将解析得到的 Book 实例放入解析结果中的 Items 字段中
	}

	fmt.Println(book) // 打印书籍信息
	return result     // 返回解析结果
}

// extraContent 用于从网页内容中提取指定正则表达式匹配到的内容
func extraContent(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
