package fetch

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	req.Header.Set("User-Agent", "crawler_book")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}
