package main

import "fmt"

func main() {
	fmt.Println(" Estruturas de controle:")

	numero := -10

	if numero > 15 {
		fmt.Println(" Maior que 15")
	} else {
		fmt.Println(" Menor ou igual a 15")
	}

	// a var do if int só é válido dentro do escopo do if
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println(" Número é maior que zero")
	} else if numero < -10 {
		fmt.Println(" Número é menor que -10")
	} else {
		fmt.Println(" Número entre 0 e -10")
	}
}
