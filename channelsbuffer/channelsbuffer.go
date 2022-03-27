package main

import "fmt"

func main() {
	// capacidade 2
	canal := make(chan string, 2)

	canal <- "Buffer channel 1"
	canal <- "Buffer channel 2"

	mensagemUm := <-canal
	mensagemDois := <-canal

	fmt.Println(mensagemUm)
	fmt.Println(mensagemDois)

}
