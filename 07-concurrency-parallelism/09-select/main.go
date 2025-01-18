package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

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

func getFastestSite(url1, url2, url3 string) string {
	c1 := findPageTitles(url1)
	c2 := findPageTitles(url2)
	c3 := findPageTitles(url3)

	// estrutura específica para trabalhar com concorrência
	// primeiro canal que chegar vai executar o case
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	}
}

func main() {
	winner := getFastestSite(
		"https://www.gov.br/pt-br",
		"https://www.youtube.com/",
		"https://www.wikipedia.org/",
	)

	fmt.Println(winner)
}
