package main

import "fmt"

func rotina(c chan int) {
	c <- 1 // operação bloqueante, recebendo valor
	fmt.Println("Só depois que foi lido")
}

func main() {

	c := make(chan int) // canal sem buffer

	go rotina(c)
	fmt.Println(<-c) // operação bloqueante, lendo valor
	fmt.Println("Foi lido")
	fmt.Println(<-c)   // gera deadlock, c fica aguardando valor mas todas as goroutines já foram executadas e não tem mais como ser enviado valor
	fmt.Println("Fim") // não será executado
}
