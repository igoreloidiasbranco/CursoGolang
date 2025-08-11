package main

import (
	"fmt"
)

func main() {

	s := make([]int, 10) // cria um slice através do método make, do tipo inteiro com 10 posições

	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)        // esse terceiro parâmetro 20, é para criar um array desse slice com 20 posições
	fmt.Println(s, len(s), cap(s)) // cap - capacidade máxima desse slice, tamanho do array interno que esse slice referência "20"

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // qdo adiciona elemento a mais do que o tamanho do array interno, automáticamente é reajustado o tamanho do array interno
	fmt.Println(s, len(s), cap(s))
}
