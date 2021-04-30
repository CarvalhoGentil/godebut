package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	Cachacas = []Cachaca{
		{Nome: "51", Volume: "974ml", Custo: "8"},
		{Nome: "Matuta", Volume: "1000ml", Custo: "25"},
	}
	debut1()
}

// pageInitial (Home Page) função para mostrar uma home page qualquer
func pageInitial(response http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpont: homePage")

	fmt.Fprint(response, "2ez4Flz")
	fmt.Println("Bem vindo ao Go Web API")
}

// aProposDe (About) é para mostrar minhas informações
func aProposDe(response http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpont: about")

	qui := "Felix Neto"

	fmt.Fprint(response, "A propos de...", qui)
	fmt.Println("Meu nome: ", qui)
}

// listaCachacas é o endpoint para listar todas as cachaças cadastradas
func listerToutesCachacas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: listerToutesCachacas")

	json.NewEncoder(w).Encode(Cachacas)
}

// debut1 é a função que vai chamar as paginas
func debut1() {

	// roteur := mux.NewRouter().StrictSlash(true)

	http.HandleFunc("/", pageInitial)
	http.HandleFunc("/aproposde", aProposDe)
	http.HandleFunc("/listertoutescachacas", listerToutesCachacas)
	// roteur.HandleFunc("/", pageInitial)
	// roteur.HandleFunc("/aproposde", aProposDe)
	// roteur.HandleFunc("/listertoutescachacas", listerToutesCachacas)

	log.Fatal(http.ListenAndServe(":8085", nil))
}

// Cachaca é a estrutura base para o ojeto cachaça
type Cachaca struct {
	Nome   string `json:"Nome"`
	Volume string `volume:"Volume"`
	Custo  string `json:"Custo"`
}

var Cachacas []Cachaca
