package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	//i := 0

	//for i < 10 {
	//	i++
	//	fmt.Println("Incrementando 1")
	//	time.Sleep(time.Second)
	//}
	//fmt.Println(i)

	//for j := 0; j < 10; j += 2 {
	//	fmt.Println("Incremento j", j)
	//	time.Sleep(time.Second)
	//}

	//nomes := []string{"Jonas", "Victor", "Yan"}

	//for _, nome := range nomes {
	//	fmt.Println(nome)
	//	time.Sleep(time.Second)
	//}

	//for indice, letra := range "PALAVRA" {
	//	fmt.Println(indice, string(letra))
	//	time.Sleep(time.Second)
	//}

	usuario := map[string]string{
		"nome":      "Jonas",
		"sobrenome": "Victor",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

}
