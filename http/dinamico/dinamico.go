package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// função que recebe dois parâmetros referentes a resposta(ResponseWriter) e requisição(Request)
func horaCerta(w http.ResponseWriter, r *http.Request) {
	//Format é uma função que formata a hora atual
	s := time.Now().Format("02/01/2006 15:04:05")

	//Fprintf é uma função que escreve no ResponseWriter
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
