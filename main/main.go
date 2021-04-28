package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	debut1()
}

// pageInitial (Home Page) função para mostrar uma home page qualquer
func pageInitial(response http.ResponseWriter, r *http.Request) {
	fmt.Fprint(response, "2ez4Flz")
	fmt.Println("Bem vindo ao Go Web API")
	fmt.Println("Endpont: homePage")
}

// aProposDe (About) é para mostrar minhas informações
func aProposDe(response http.ResponseWriter, r *http.Request) {
	qui := "Felix Neto"

	fmt.Fprint(response, "A propos de...", qui)
	fmt.Println("Meu nome: ", qui)
	fmt.Println("Endpont: about")
}

// debut1 é a função que vai chamar as paginas
func debut1() {

	http.HandleFunc("/", pageInitial)
	http.HandleFunc("/about", aProposDe)

	log.Fatal(http.ListenAndServe(":8085", nil))
}
