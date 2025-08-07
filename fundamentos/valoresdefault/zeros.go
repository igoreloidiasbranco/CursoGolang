package main

import "fmt"

func main() {

	var a int
	var b float64
	var c bool
	var d string
	var e *int // * significa que é um ponteiro (nesse caso ponteiro de um inteiro)

	fmt.Printf("%v %v %v %v %v \n", a, b, c, d, e) //(o string por padrão é vazio, na saída será mostrado um espaço em branco(vazio)) e <nil> quer dizer nulo.

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e) //(substituindo por %q, a string vazia é mostrada (""))
}
