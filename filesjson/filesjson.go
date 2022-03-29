package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {
	// struct para JSON
	cUm := cachorro{"Rex", "Dalmata", 3}

	cachorroUmEmJSON, erro := json.Marshal(cUm)
	if erro != nil {
		log.Fatal(erro)
	}

	// Aqui ele printar um slice de bytes ...
	fmt.Println(cachorroUmEmJSON)
	// por isso é necessário printar com o bytes para conversao
	fmt.Println(bytes.NewBuffer(cachorroUmEmJSON))

	// outro exemplo agora com map para JSON
	cDois := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorroDoisEmJSON, erro := json.Marshal(cDois)
	if erro != nil {
		log.Fatal(erro)
	}

	// Aqui ele printar um slice de bytes ...
	fmt.Println(cachorroDoisEmJSON)
	// por isso é necessário printar com o bytes para conversao
	fmt.Println(bytes.NewBuffer(cachorroDoisEmJSON))
}
