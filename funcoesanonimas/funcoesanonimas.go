package main

import "fmt"

func main() {

	func(texto string) {
		fmt.Println(texto)
	}("Passando Parâmetro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Mais sobre anonimas")

	fmt.Println(retorno)

}
