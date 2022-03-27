package main

import "fmt"

// Padroes de Concorrencia: Workers Pools
func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// no meu raspbarry pi 4, usei 4 goroutines e usei o comando
	// time go run .
	// go run .  47,45s user 1,74s system 155% cpu 31,585 total
	// ao testar com apens 1 goroutines logo depois de usar 4:
	// go run .  48,25s user 1,16s system 101% cpu 48,898 total
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)

	}
}

// no caso tarefas só <-recebe dados e resultados só <-envia dados
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

// 1 1 2 3 5 8 13
func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
