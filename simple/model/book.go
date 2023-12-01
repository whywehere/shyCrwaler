package model

// Book 是表示书籍的结构体，包含了书名、作者、出版社、页数、价格、评分和简介等信息
type Book struct {
	Name      string  // 书名
	Author    string  // 作者
	Publisher string  // 出版社
	Pages     int     // 页数
	Price     string  // 价格
	Score     float64 // 评分
	Intro     string  // 简介
}
