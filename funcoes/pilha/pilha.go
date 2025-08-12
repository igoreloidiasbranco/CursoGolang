package main

import "runtime/debug"

func f3() {
	debug.PrintStack() //pacote debug, imprimi a pilha de execução do programa no momento que a função foi chamada.
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}
