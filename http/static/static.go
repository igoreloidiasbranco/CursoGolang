//Servidor para servir arquivos estáticos

package main

import (
	"log"
	"net/http"
)

func main() {
	//FileServer da biblioteca http.FileServer e o parâmetro é o diretório que será lido
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
