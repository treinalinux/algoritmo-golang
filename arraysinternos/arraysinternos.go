package main

import "fmt"

func main() {
	// make está criando um array de 15 posicoes
	// e retorna um slice de 10 primerias posicoes
	sliceUm := make([]float32, 10, 15)
	fmt.Println("-------------------------")
	fmt.Println(sliceUm)
	// Tamanho
	fmt.Println(" Tamanho: ", len(sliceUm))
	// Capacidade
	fmt.Println(" Capacidade: ", cap(sliceUm))
	fmt.Println("-------------------------")

	sliceUm = append(sliceUm, 5)
	fmt.Println(sliceUm)
	fmt.Println(" Tamanho: ", len(sliceUm))
	fmt.Println(" Capacidade: ", cap(sliceUm))
	fmt.Println("-------------------------")

	sliceUm = append(sliceUm, 5)
	fmt.Println(sliceUm)
	fmt.Println(" Tamanho: ", len(sliceUm))
	fmt.Println(" Capacidade: ", cap(sliceUm))
	fmt.Println("-------------------------")

	sliceUm = append(sliceUm, 5)
	fmt.Println(sliceUm)
	fmt.Println(" Tamanho: ", len(sliceUm))
	fmt.Println(" Capacidade: ", cap(sliceUm))
	fmt.Println("-------------------------")

	sliceUm = append(sliceUm, 5)
	fmt.Println(sliceUm)
	fmt.Println(" Tamanho: ", len(sliceUm))
	fmt.Println(" Capacidade: ", cap(sliceUm))
	fmt.Println("-------------------------")

	sliceUm = append(sliceUm, 5)
	fmt.Println(sliceUm)
	fmt.Println(" Tamanho: ", len(sliceUm))
	fmt.Println(" Capacidade: ", cap(sliceUm))
	fmt.Println("-------------------------")

	// agora o go vê que pode estourar a " Capacidade: ", capcidade
	// entao ele dobra
	sliceUm = append(sliceUm, 5)
	fmt.Println(sliceUm)
	fmt.Println(" Tamanho: ", len(sliceUm))
	fmt.Println(" Capacidade: ", cap(sliceUm))
	fmt.Println("-------------------------")
}
