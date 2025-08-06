package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.149516

	// mesclar float64 com String não funciona em Go com fmt
	//fmt.Println("O valor de x é: " + x)

	//formas de concatenar

	//criar uma nova versão String da variável x usando uma função Sprint
	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é: " + xs)

	//outra maneira sem utilizar função Sprint
	fmt.Println("O valor de x é:", x, "!!!")

	//ou usar função printf, que tem mais funcionalidades, como por exemplo diminuir as casas decimais (obs: faz arredondamento)
	fmt.Printf("O valor de x é: %.2f", x)

	a, b, c, d := 1, 1.999, false, "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	//%v tbm serve para quase todos os valores
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
