package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nome string
	var idade int
	fmt.Print("Informe, nome & idade: ")
	fmt.Scanf("%s %d", &nome, &idade)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Informe seu cargo: ")

	cargo, _ := reader.ReadString('\n')

	fmt.Printf(" Olá, %v, %v está cheio de vida com %v anos... \n", nome, cargo, idade)

}
