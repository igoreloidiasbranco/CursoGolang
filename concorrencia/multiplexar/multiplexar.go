package main

//padrão de concorrência
//multiplexar - misturar (mensagens) num canal
import (
	"curso-golang/html"
	"fmt"
)

// primeiro channel <-chan é somente leitura, segundo channel é para envio de dados
func encaminhar(origem <-chan string, destino chan string) {

	for {
		//recebendo dados do canal de origem e enviando para o canal de destino
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(

		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.amazon.com.br", "https://www.youtube.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
