package main

import "fmt"

func main() {

	x := 1
	y := 2

	// operadores unários

	//incremento
	x++

	//decremento
	y--

	fmt.Println(x, y)

	//fmt.Println(++x == y--) Em Go não é permitido comparação unido com incrementos ou decrementos.
}
