package main

import (
	"fmt"
)

// Assinatura da função
func somar(a int, b int) int { // sintaxe: func (pra criar a função) + nome da função("somar") + () parâmetros(nome da variável e tipo) + retorno(int)
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}
