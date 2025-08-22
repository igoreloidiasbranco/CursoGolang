package main

import (
	"fmt"
	"runtime"
)

//Concorrência vs Paralelismo
//Concorrência: é quando temos múltiplos processos ou threads executando ao mesmo tempo.
//Paralelismo: é quando temos múltiplos processos ou threads executando ao mesmo tempo, mas em diferentes núcleos do processador.

func main() {

	fmt.Println(runtime.NumCPU)
}
