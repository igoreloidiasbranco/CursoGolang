package main

import "fmt"

func main() {

	//Map, lista chave-valor

	//var aprovados map[int]string
	//mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98765432198] = "Pedro"
	aprovados[98767865423] = "Ana"

	fmt.Println(aprovados)

	//percorre a lista
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	//pegar valor, passando a chave
	fmt.Println(aprovados[98765432198])

	//deletar - passa o map e a chave
	delete(aprovados, 98765432198)

	fmt.Println(aprovados)
}
