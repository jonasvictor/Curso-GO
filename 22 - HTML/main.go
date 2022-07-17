package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{"Jonas", "jonas.victor@gmail.com"}
	templates.ExecuteTemplate(w, "home.html", u)
}

var templates *template.Template

func main() {

	templates = template.Must((template.ParseGlob("*.html")))

	http.HandleFunc("/home", home)

	fmt.Println("Executando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
