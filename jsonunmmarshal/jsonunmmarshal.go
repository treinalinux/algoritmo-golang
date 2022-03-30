package main

import (
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
	cachorroUmEmJSON := `{"nome": "Rex", "raca": "Dalmata", "idade": 3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroUmEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorroDoisEmJSON := `{"nome": "Tom", "raca": "Poodle"}`

	cachorroDois := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorroDoisEmJSON), &cachorroDois); erro != nil {
		log.Fatal(erro)

	}
	fmt.Println(cachorroDois)

}
