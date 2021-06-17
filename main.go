package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Iniciando alambique em http://localhost:8085 ")
	Cachacas = []Cachaca{
		{Nome: "51", Volume: "974ml", Custo: "8"},
		{Nome: "Matuta", Volume: "1000ml", Custo: "25"},
	}
	debut1()
}

// pageInitial (Home Page) função para mostrar uma home page qualquer
func pageInitial(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpont: homePage")

	fmt.Fprintln(w, "-- 2ez4Flz --")
	fmt.Fprint(w, "Go Drink")
}

// aProposDe (About) é para mostrar minhas informações
func aProposDe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpont: about")

	qui := "Felix Neto"

	fmt.Fprint(w, "A propos de...", qui)
}

// listerToutesCachacas é o endpoint para listar todas as cachaças cadastradas
func listerToutesCachacas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: listerToutesCachacas")

	json.NewEncoder(w).Encode(Cachacas)
}

// debut1 é a função que vai chamar as paginas
func debut1() {

	roteur := mux.NewRouter().StrictSlash(true)

	roteur.HandleFunc("/", pageInitial)
	roteur.HandleFunc("/aproposde", aProposDe)
	roteur.HandleFunc("/listertoutescachacas", listerToutesCachacas)

	log.Fatal(http.ListenAndServe(":8085", roteur))
}

// Cachaca é a estrutura base para o ojeto cachaça
type Cachaca struct {
	Nome   string `json:"Nome"`
	Volume string `volume:"Volume"`
	Custo  string `json:"Custo"`
}

var Cachacas []Cachaca
