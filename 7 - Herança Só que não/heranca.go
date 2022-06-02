package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	pessoa1 := pessoa{"Jonas", "Victor", 25, 1.73}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Ciência da Computação", "UFPB"}
	fmt.Println(estudante1)
}
