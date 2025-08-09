package main

import (
	"fmt"
	"math/rand"
	"time"
)

//em Go, é possível ter um bloco de inicialização antes de executar a estrutura e esse bloco fornece uma variável local dentro do scopo if ou else

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) //rand, pacote do Go(randow), time.Now().UnixNano() -> pega nano segungo da data atual e passa como fonte pro rand gerar o número aleatório
	r := rand.New(s)

	return r.Intn(10) // gera o número aleatório até 10
}

func main() {
	if i := numeroAleatorio(); i > 5 { // tbm suportado no switch
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!!!")
	}
}
