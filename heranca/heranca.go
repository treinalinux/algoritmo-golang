package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	pessoaUm := pessoa{"Alan", "Alves", 29, 190}
	fmt.Println(pessoaUm)

	estudanteUm := estudante{pessoaUm, "Ciência da Computacão", "Faculdade Descomplica"}
	fmt.Println(estudanteUm)
	fmt.Println(estudanteUm.curso)
	fmt.Println(estudanteUm.faculdade)
	fmt.Println(estudanteUm.sobrenome)
}
