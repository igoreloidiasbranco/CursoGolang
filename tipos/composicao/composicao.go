package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

// interface com composição de outra interface
type esportivoLuxuoso interface {
	esportivo
	luxuoso
	//pode adicionar novos métodos
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Ligando Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Fazendo Baliza...")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
