package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Jonas",
		"sobrenome": "Victor",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Jonas",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome": "LCC",
			"UFPB": "CAMPUS-IV",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Peixes",
	}

	fmt.Println(usuario2)
}
