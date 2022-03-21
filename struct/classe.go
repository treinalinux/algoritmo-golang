package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Em Go uma classe é na verdade um struct")

	// enderecoUm := endereco{"Rua nova", 0}

	var userUm usuario
	userUm.nome = "Alan"
	userUm.idade = 98
	fmt.Println(userUm)

	userDois := usuario{"Carla", 44}
	fmt.Println(userDois)

	userTres := usuario{nome: "Fernanda"}
	fmt.Println(userTres)
}
