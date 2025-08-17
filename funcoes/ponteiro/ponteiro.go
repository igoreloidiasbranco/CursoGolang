package main

import "fmt"

func inc1(n int) {

	n++
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	//revisão: * é usado para desreferenciar, ou seja, ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {

	n := 1
	inc1(n)        //passando uma CÓPIA do valor de n e a função inc1 opera em cima da cópia, o que não afeta no valor do n original
	fmt.Println(n) //(A saída aqui será 1)

	// revisão: & usado para obter o endereço da variável
	inc2(&n)       // passando o endereço da referência, nesse caso o valor do n original será modificado
	fmt.Println(n) //(A saída aqui será 2)

	//priorize o envio por cópia, para evitar efeito colateral que a alteração do valor real possa afetar em outros métodos que utilizam essa variável
}
