package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// função trata das goroutines e retorna um channel para consumo
// padrão generator
func findPageTitles(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	c := findPageTitles("https://www.gov.br/pt-br", "https://www.youtube.com/")

	fmt.Println(<-c)
}
