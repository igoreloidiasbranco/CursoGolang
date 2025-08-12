package main

import "fmt"

// função anônima, não tem nome, mas atribui valor a variável soma
var soma = func(a, b int) int {
	return a + b
}

func main() {

	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
