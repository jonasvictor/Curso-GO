package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int64 = -1000000000000000000
	fmt.Println(numero)

	var numero2 uint64 = 1000000000000000000
	fmt.Println(numero2)

	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1230000.45
	fmt.Println(numeroReal3)

	// Fim numeros reais

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'J'
	fmt.Println(char)

	// Fim strings

	texto := 5
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
