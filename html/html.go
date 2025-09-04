package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// Titulo obtem o título de uma página html
func Titulo(urls ...string) <-chan string {
	c := make(chan string, len(urls))
	for _, url := range urls { //for_ ignora o índice
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			matches := r.FindStringSubmatch(string(html))
			if len(matches) > 1 {
				c <- matches[1]
			} else {
				c <- "Título não encontrado"
			}
		}(url)

	}
	return c
}
