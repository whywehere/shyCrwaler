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
	authorRe    = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
	publisherRe = regexp.MustCompile(`<span class="pl">出版社:</span>[\d\D]*?<a.*?>([^<]+)</a>`)
	pagesRe     = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
	priceRe     = regexp.MustCompile(`<span class="pl">定价:</span> ([^<]+)<br/>`)
	scoreRe     = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+) </strong>`)
	introRe     = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)
)

func ParseBookDetail(contents []byte, name string) engine.ParseResult {
	book := model.Book{}
	book.Name = name
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
		Items:    []interface{}{book},
	}
	fmt.Println(book)
	return result
}

func extraContent(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
