package main

import (
	"fmt"
)

func main() {

	// certificando que dois slices apontam para o mesmo array interno
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[0] = 7
	fmt.Println(s1, s2) // aqui podemos verificar que tanto para s1 e s2, o index 0 foi alterado para o valor 7, que significa que apontaram para o mesmo array interno.
}
