package main

import (
	"fmt"
	"time"
)

// CONCORRÊNCIA != PARALELISMO
func main() {
	fmt.Println("Concorrência")
	go escrever("Olá Mundo!") //goroutines
	escrever("programando em GO!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)

	}
}
