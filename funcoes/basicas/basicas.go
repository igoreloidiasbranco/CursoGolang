package main

import "fmt"

//Sintaxe  func + nome da função + parâmetros + retorno

// função sem parâmetro e sem retorno
func f1() {
	fmt.Println("Função F1")
}

// função com parâmetro (nome + tipo, separados por vírgula) sem retorno
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

// função sem parâmetros com retorno
func f3() string {
	return "F3"
}

// função com parâmetros de forma optimizada com retorno
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2) //Sprintf não imprime no console, mas retorna uma string formatada
}

// função em Go pode retornar multiplos valores

// função sem parâmetros e que retorna dois valores
func f5() (string, string) {
	return "retorno 1", "retorno 2"
}

func main() {
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")

	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)

	r51, _ = f5()
	fmt.Println("F5:", r51)
}
