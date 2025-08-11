package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Slice - maior flexibilidade no uso dos arrays
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slide define um pedaço de um array. (trecho de um array), mas que não gera um array diferente

	s2 := a2[1:3] // slice s2 atribui o index de um a 3 de a2, mas não incluindo o index 3, portanto index 1 e 2 apenas
	fmt.Println(a2, s2)

	s3 := a2[:2] // slice s3 atribui o ínicio do index a2, até o index 2, sem incluir o index 2.
	fmt.Println(a2, s3)

	// vc pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1] // s4 slice do slice s2
	fmt.Println(s2, s4)
}
