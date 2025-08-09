package main

import (
	"fmt"
	"time"
)

func main() {

	// em Go, a única estrutura com controle de repetição é o for, não temos uma estrutura como while por exemplo.

	// usar um for como se fosse um while
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	for {
		//laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
		// ctrl + alt + m para parar no code runner
	}
}
