package fetch

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Fetch 用于获取指定URL的网页内容
func Fetch(url string) ([]byte, error) {
	// 创建GET请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	// 设置请求头
	req.Header.Set("User-Agent", "crawler_book")

	// 发送请求并获取响应
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// 读取响应内容并返回
	return io.ReadAll(resp.Body)
}
