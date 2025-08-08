package main

import "fmt"

func main() {

	i := 1

	// Ponteiro é uma referência de um espaço alocado em memória (*) que em seguida, determina o tipo da varíavel do ponteiro

	var p *int = nil

	// (&)esse comando pega o endereço da variável i e atribui a variável p
	p = &i

	*p++ //pega o valor do endereço de memória que o ponteiro está apontando e atribui valor.
	i++  // atribui valor no mesmo endereço de memória que o ponteiro de p está apontando.

	fmt.Println(p, *p, i, &i)

	// Go não tem aritmética em cima de ponteiros. Ex: p++ não é permitido
}
