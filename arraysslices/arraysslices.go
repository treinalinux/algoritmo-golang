package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var arrayUm [5]string
	arrayUm[0] = "Posicao 1"
	fmt.Println(arrayUm)

	arrayDois := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arrayDois)

	sliceUm := []int{}
	fmt.Println(sliceUm)

	sliceDois := []int{10, 20, 30, 40, 50}
	fmt.Println(sliceDois)

	// TypeOf devolve o tipo de uma var
	fmt.Println(reflect.TypeOf(sliceUm))
	fmt.Println(reflect.TypeOf(arrayDois))

	sliceUm = append(sliceUm, 90)
	fmt.Println(sliceUm)

	sliceUm = append(sliceUm, 80, 70)
	fmt.Println(sliceUm)

	sliceTres := arrayDois[1:3]
	fmt.Println(sliceTres)
}
