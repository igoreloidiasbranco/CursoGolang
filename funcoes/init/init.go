package main

import "fmt"

// INIT - convenção usada em Go, função executada independente de ser acionada
func init() {

	fmt.Println("Inicializando...") //esse trecho será processado antes da execução main
}

func main() {

	fmt.Println("Main...")
}
