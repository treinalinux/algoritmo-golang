package main

import "fmt"

func main() {

	// map aqui é chave do tipo string e valor do tipo string
	usuario := map[string]string{
		"nome":      "Alan",
		"sobrenome": "Alves",
	}

	// map aqui é chave do tipo string e valor do tipo int
	numero := map[string]int{
		"telefone": 8888888888,
	}

	// map aqui é chave do tipo string e valor do tipo map
	faculdade := map[string]map[string]string{
		"bachareladoTecUm": {
			"curso": "Ciencia da Computacao",
			"tempo": "4 anos",
		},
		"bachareladoTecDois": {
			"curso": "Sistema da Informcao",
			"tempo": "4 anos",
		},
	}

	fmt.Println(usuario["nome"])
	fmt.Println(numero["telefone"])

	fmt.Println(faculdade)
	delete(faculdade, "bachareladoTecDois")
	fmt.Println(faculdade)

}
