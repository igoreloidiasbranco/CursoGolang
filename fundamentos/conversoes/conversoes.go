package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2

	fmt.Println(x / float64(y)) // converte int y em float64

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal) // convert float64 em int, arredonda pra baixo (resultado é 6)

	// CUIDADO conversao de int para string, dessa maneira o resultado é o da tabela ask (será a letra 'a')
	fmt.Println("Teste " + string(97))

	//conversão correta
	fmt.Println("Teste " + strconv.Itoa(97)) // strcon.Itoa (converte inteiro para string)

	// string para int
	num, _ := strconv.Atoi("97") //em Go, pode-se retornar dois valores na chamada de uma função, nesse caso o num é o retorno de um número e o _ é o retorno de um erro caso exista.
	fmt.Println(num - 96)

	// string para boolean
	b, _ := strconv.ParseBool("true") // true ou o valor 1, se colocar 0 ele entende como false
	if b {
		fmt.Println("Verdadeiro")
	}
}
