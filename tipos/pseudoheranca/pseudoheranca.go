package main

import (
	"fmt"
)

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // (é uma composição) em Go, não tem herança
	turboLigado bool
}

func main() {

	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f) // a saída aqui será {{F40 0} true} ou seja, não é uma herança, temos 2 estruturas, 1-carro e 1-ferrari
}
