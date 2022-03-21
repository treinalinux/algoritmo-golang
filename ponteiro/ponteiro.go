package main

import "fmt"

func main() {
	var variavelUm int = 10
	var variavelDois int = variavelUm

	fmt.Println(variavelUm, variavelDois)

	variavelUm++
	fmt.Println(variavelUm, variavelDois)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavelTres int // aqui é um inteiro vai para memória
	var ponteiroUm *int  // aqui é uma referência de um inteiro que vai para memória

	variavelTres = 100
	// assim está errado:
	// ponteiroUm = variavelTres
	// cannot use variavelTres (type int) as type *int in assignment
	ponteiroUm = &variavelTres

	// desferenciacao lê o valor dentro do *ponteiroUm
	fmt.Println(variavelTres, *ponteiroUm)
}
