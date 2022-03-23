package main

// Fonte:
// https://zetcode.com/golang/readinput/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	names := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Informe seu nome: ")

		scanner.Scan()

		text := scanner.Text()

		if len(text) != 0 {

			fmt.Println(text)
			names = append(names, text)
		} else {
			break
		}
	}

	fmt.Println(names)

}
