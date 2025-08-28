//canal sem buffer - qdo envia um dado para o canal, aquela linha de código vai ficar esperando até que o dado seja consumido
//canal com buffer - consegue enviar vários dados para o canal, até que o buffer fique cheio gerando o bloqueio, até que os dados
//sejam consumidos

package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {

	ch <- 1                  //esse é o que será lido na chamada da main
	ch <- 2                  //add no buffer
	ch <- 3                  //add no buffer
	ch <- 4                  //buffer cheio
	ch <- 5                  //fica aguardando até liberar espaço no buffer
	fmt.Println("Executou!") //aqui não será executado pois o buffer já está no limite
	ch <- 6
}

func main() {

	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
