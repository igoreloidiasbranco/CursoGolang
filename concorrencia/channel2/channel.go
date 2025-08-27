// Channel (canal) é a forma de comunicação entre goroutines, ponto de sincronismo, aguarda o dado chegar
// Goroutine roda de forma independente e o canal é o ponto de parada das goroutines
package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int) {

	time.Sleep(time.Second)
	c <- 2 * base //enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c //recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}
