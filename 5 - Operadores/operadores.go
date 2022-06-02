package main

import "fmt"

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 1 * 2
	divisao := 1 / 2
	resto := 1 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// Fim dos aritméticos

	// Atribuição

	var variavel string = "String"
	variavel2 := "String2"
	fmt.Println(variavel, variavel2)

	// Operadores relacionais

	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 < 2)
	fmt.Println(1 > 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 >= 2)

	// Operadores lógicos

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Fim dos ooperadores lógicos

	// Operadores unários

	fmt.Println("-----------------------------")
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 15 // numero = numero - 15
	fmt.Println(numero)

	numero *= 2 // numero = numero * 2
	fmt.Println(numero)

	numero /= 2 // numero = numero / 2
	fmt.Println(numero)

	numero %= 2 // numero = numero % 2
	fmt.Println(numero)

	fmt.Println(numero)
	fmt.Println("-----------------------------")

	// Fim dos operadores unários

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
