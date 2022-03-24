package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func main() {
	totalDaSoma := soma(10, 20, 30, 40)
	fmt.Println(totalDaSoma)

}
