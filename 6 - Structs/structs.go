package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arqiovo Structs")

	var u usuario
	u.nome = "Jonas"
	u.idade = 25
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua principal", 0}

	usuario2 := usuario{"Jonas", 25, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Jonas", idade: 25}
	fmt.Println(usuario3)
}
