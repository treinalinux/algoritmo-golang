package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 5 {
		i++
		fmt.Printf("Incrementando o i: %v\n", i)
		time.Sleep(time.Second)
	}

	for j := 0; j <= 10; j += 2 {
		fmt.Printf("Incrementando o j: %v\n", j)
		time.Sleep(time.Second)
	}

	nomes := [4]string{"Alan", "da", "Silva", "Alves"}

	// quiser omitir o indice é necessário usar o "_"
	for indice, nome := range nomes {
		fmt.Printf("-- %v: %v\n", indice, nome)
	}

	for indice, letra := range "LETRAS" {
		fmt.Printf("-- Indice: %v Tabela ASCII: %v Letra: %v\n", indice, letra, string(letra))
	}

	curriculo := map[string]string{
		"nome":          "Alan",
		"sobrenome":     "Alves",
		"desenvolvedor": "Go",
		"graduacao":     "Ciência da Computacao",
	}

	for chave, valor := range curriculo {
		fmt.Println(chave, valor)
	}

	//for {
	//	fmt.Println("Loop Infinito")
	//}
}
