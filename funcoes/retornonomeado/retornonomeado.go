package main

import "fmt"

// return nomeado

// função que inverte a ordem dos parâmetros e retorna sem a necessidade de explicitar o return pois já foi sinalizado nos parâmetros de retorno o tipo e a ordem.
func trocar(p1, p2 int) (segundo int, primeiro int) {

	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {

	x, y := trocar(2, 3)
	fmt.Println(x, y)
}
