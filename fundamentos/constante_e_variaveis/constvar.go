package main

import (
	"fmt"
	"math"
)

func main() {

	// sintaxe de constante const + nome da constante + tipo = valor
	const PI float64 = 3.1415

	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de declarar uma variável com inicialização (:= declara e inicializa) é a forma mais utilizada
	// o Go não compila se a variável declarada não for usada
	area := PI * math.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)

	//outra forma de declarar constantes
	const (
		a = 1
		b = 2
	)

	// outra forma de declarar variáveis
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// pode declarar duas ou mais variáveis em uma única linha. var e, f bool = true, false , true vai para a variável e e false para a variável f
	var e, f bool = true, false

	fmt.Println(e, f)

	// tbm de forma reduzida
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
