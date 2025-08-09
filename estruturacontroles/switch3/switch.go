package main

import (
	"fmt"
	"time"
)

// switch em cima de parâmetros genéricos, conforme o parâmetro passado, entra em algum case.

func tipo(i interface{}) string { //i interface recebe parâmetro genérico de qualquer tipo sem erro de compilação

	switch i.(type) {

	case int:
		return "Inteiro"

	case float32, float64:
		return "Valor flutuante"

	case string:
		return "string"

	case func():
		return "função"

	default:
		return "Não sei"
	}
}

func main() {

	fmt.Println(tipo(1))
	fmt.Println(tipo(3.2))
	fmt.Println(tipo("Tipo string"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
