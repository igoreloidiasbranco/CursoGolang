package main

import "fmt"

//DEFER - Função reservada da linguagem Go, que atrasa a execução de uma sequência de código até o momento antes da chamada do return
//Ex: fechar algum tipo de recurso que está aberto antes da saída do método.

func obterValorAprovado(numero int) int {

	defer fmt.Println("fim!") //isso só será executado momento antes de sair do método, momento antes do return.
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} else {
		fmt.Println("Valor baixo...")
		return numero
	}
}

func main() {

	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
