package main

import "fmt"

func main() {

	// for com range, equivalente ao for each.

	// Array com tamanho fixado conforme a quantidade de elementos
	numeros := [...]int{1, 2, 3, 4, 5}

	for i, numero := range numeros { // esse for range retorna o index e o valor do array numeros
		fmt.Printf("%d) %d\n", i, numero)
	}

	// para ignorar o valor do index, substitua i por _
	for _, num := range numeros {
		fmt.Println(num)
	}

}
