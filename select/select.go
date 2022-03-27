package main

import (
	"fmt"
	"time"
)

func main() {
	canalUm, canalDois := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Microsecond * 500)
			canalUm <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Microsecond * 500)
			canalDois <- "Channel 2"
		}
	}()

	for {

		select {
		case mensagemCanalUm := <-canalUm:
			fmt.Println(mensagemCanalUm)
		case mensagemCanalDois := <-canalDois:
			fmt.Println(mensagemCanalDois)
		}
	}
}
