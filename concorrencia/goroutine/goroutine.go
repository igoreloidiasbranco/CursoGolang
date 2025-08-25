//Goroutine fundamentos da concorrência no Go

package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {

	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// Ex de uso sequêncial, sem uso da concorrência
	//fale("Maria", "porque vc não fala comigo?", 3)
	//fale("João", "só posso falar depois de você!", 1)

	// Ex de processamentos independentes (goroutine) que podem ser executados ao mesmo tempo
	//usando o comando go antes da chamada do método, (goroutine) como se fosse uma THREAD(linha de execução independente que pode ser executada em paralelo)
	//goroutine só continua executando se a Thread principal do programa estiver funcionando,
	//nesse caso, não dará tempo da goroutine ser executada devido ao time.sleep de 1 seg, a main finalizará antes disso.
	//go fale("Maria", "Ei...", 10)
	//go fale("João", "Opa...", 10)

	//nesse caso, não dará tempo da Maria falar 10 vezes, pois João encerrará antes a THREAD principal
	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5)

	// para que a main aguarde a execução das goroutines utilizamos o CHANNEL que veremos em outro exemplo.
}
