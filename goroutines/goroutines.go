package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	// Usando o go na frente habilita a goroutines
	go escreva("Golang")
	escreva("Ruby")
}

func escreva(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
