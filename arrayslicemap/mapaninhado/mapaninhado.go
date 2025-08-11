package main

import "fmt"

func main() {

	// map aninhado
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.23,
		},

		"J": {
			"José João": 11325.43,
		},

		"P": {
			"Pedro Junior": 1230.40,
		},
	}

	delete(funcsPorLetra, "P") //deleta todos os funcionários com a letra P

	for letra, funcionario := range funcsPorLetra {
		fmt.Println(letra, funcionario)
	}
}
