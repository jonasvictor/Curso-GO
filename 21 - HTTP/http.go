package main

import (
	"log"
	"net/http"
)

func raiz(w http.ResponseWriter, r *http.Request) { // raiz
	w.Write([]byte("Página raiz!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar páginas de usuários!"))
}

func main() {

	http.HandleFunc("/", raiz)

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICCAÇÃO WEB

// CLIENTE (FAZ REQUISIÇÃO)- SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)

// Request - Response

//	 Rotas
// 		URI - Identificador do recurso
// 		Método - GET, POST, PUT, DELETE
