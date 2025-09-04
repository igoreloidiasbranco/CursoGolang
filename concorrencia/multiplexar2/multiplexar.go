package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {

	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ { // for infinito fica executando até a execução da goroutine ser encerrada
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}() // () chamada da função anônima func()
	return c
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {

	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			}
		}
	}() // () chamada da função anônima func()
	return c
}

func main() {
	c := juntar(falar("João"), falar("Maria"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
