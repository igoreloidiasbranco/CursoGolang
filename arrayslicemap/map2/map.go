package main

import "fmt"

func main() {

	//map de chave funcionário e valor salário
	funcsESalario := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Júnior":   2200.50, // o último elemento também precisa ter vírgula
	}

	//acrescentando
	funcsESalario["Rafael Filho"] = 1350.0
	delete(funcsESalario, "Inexistente") // não tem problema excluir uma chave que não existe

	for nome, salario := range funcsESalario {
		fmt.Println(nome, salario)
	}
}
