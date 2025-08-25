// Channel é um tipo da linguagem Go, assim como um inteiro, ou outro tipo.
// com o canal, é possível aguardar a execução das goroutines executadas de forma independente(assíncronas)
package main

import (
	"fmt"
)

func main() {
	// make(chan de channel(canal) e tipo + buffer(1))
	ch := make(chan int, 1)

	//agregando um valor para o channel do tipo int (escrita)
	ch <- 1

	//recebendo dados do canal (leitura) e remove o buffer valor do canal(1) ficando vazio
	<-ch

	ch <- 2 // agregando valor novamente
	fmt.Println(<-ch)
}
