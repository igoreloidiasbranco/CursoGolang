// Google I/O 2012 - Go Concurrency Patterns

// <-chann - canal somente leitura, somente recebe dados nesse canal

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Generator - é uma função que dentro dela tem uma goroutine que atende multiplas chamadas
func titulo(urls ...string) <-chan string {
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

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com.br", "https://www.youtube.com")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
