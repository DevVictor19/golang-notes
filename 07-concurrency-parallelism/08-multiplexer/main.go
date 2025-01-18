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

func redirect(origin <-chan string, destiny chan string) {
	for {
		destiny <- <-origin
	}
}

func join(inputs ...<-chan string) <-chan string {
	c := make(chan string)

	for i := 0; i < len(inputs); i++ {
		go redirect(inputs[i], c)
	}

	return c
}

func main() {
	// multiplexar: pegar dados de canais diferentes e afunilar em um único canal

	c := join(
		findPageTitles("https://www.gov.br/pt-br", "https://www.youtube.com/"),
		findPageTitles("https://www.wikipedia.org/", "https://www.github.com/"),
		findPageTitles("https://www.bbc.com/", "https://www.reddit.com/"),
	)

	// vai exibir a requisição mais rápida
	fmt.Println(<-c)
}
