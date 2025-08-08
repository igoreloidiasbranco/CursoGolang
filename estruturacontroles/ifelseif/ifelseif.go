package main

import "fmt"

func notaParaConceito(nota float64) string {

	if nota >= 9 {
		return "A"
	} else if nota >= 8 && nota < 9 {
		return "B"
	} else if nota >= 5 && nota < 8 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {

	fmt.Println(notaParaConceito(11.8))
	fmt.Println(notaParaConceito(8.7))
	fmt.Println(notaParaConceito(6.5))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(2.7))

}
