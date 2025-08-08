package main

import "fmt"

func obterResultado(nota float64) string {

	// em Go, NÃO TEM OPERADOR TERNÁRIO
	// ex: return nota >= 6 ? "Aprovado" : "Reprovado" isso não funciona, precisa fazer com if

	if nota > 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {

	fmt.Println(obterResultado(6.2))

}
