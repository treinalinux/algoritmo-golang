package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escreva("Golang")
		waitGroup.Done()
	}()

	go func() {
		escreva("Ruby")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escreva(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
