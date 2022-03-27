package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Escrever..."), escrever("Multiplexar"))

	for i := 0; i < 30; i++ {
		fmt.Println(<-canal)

	}
}

// Multiplexador
func multiplexar(canalDeEntradaUm, canalDeEntradaDois <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalDeEntradaUm:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntradaDois:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf(" Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
