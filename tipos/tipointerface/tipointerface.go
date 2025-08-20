package main

import "fmt"

type curso struct {
	nome string
}

func main() {

	// tipo interface vazia, deixa a variável mais genérica, podendo receber outros tipos de valores, como se fosse o Object do Java

	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	coisa = true
	fmt.Println(coisa)

	type dinamico interface{}
	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a linguagem do Google"}
	fmt.Println(coisa2)
}
