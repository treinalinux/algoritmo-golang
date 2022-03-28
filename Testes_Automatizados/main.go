package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TiposDeEnderecos("Avenida Paulista")
	fmt.Println(tipoEndereco)

}
