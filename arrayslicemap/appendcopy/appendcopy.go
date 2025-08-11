package main

import (
	"fmt"
)

func main() {

	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// slice1 = append(array1, 4, 5, 6) append não funciona com array

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // copia o slice1 para o slice2, apenas os dois primeiros números do slice1 pois, slice2 tem somente 2 posições.
	fmt.Println(slice2)
}
