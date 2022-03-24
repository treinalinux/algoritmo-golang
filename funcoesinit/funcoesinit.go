package main

import "fmt"

func init() {
	fmt.Println("Funcao init... Sempre executa antes da main")
	fmt.Println("Funcao init... Pode ter uma por arquivo")

}

func main() {
	fmt.Println("Funcao main")
}
