//É possível fazer um for dentro de um canal

package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c) // comando para fechar o laço for para que o canal não fique esperando novos dados que não vão chegar(deadlock).
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c) //função cap para verificar a capacidade do canal, tamanho do buffer do canal
	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!")
}
