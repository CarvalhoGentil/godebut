package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Iniciando alambique em http://localhost:8085 ")
	Cachacas = []Cachaca{
		Cachaca{Nome: "51", Volume: "974ml", Custo: "8"},
		Cachaca{Nome: "Matuta", Volume: "1000ml", Custo: "30"},
	}
	debut1()
}

// pageInitial (Home Page) função para mostrar uma home page qualquer
func pageInitial(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: pageInitial")

	fmt.Fprintln(w, "-- 2ez4Flz --")
	fmt.Fprint(w, "-  GoDrink  -")
	fmt.Fprint(w, "- - - - - - -")
}

// aProposDe (About) é para mostrar minhas informações
func aProposDe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: aproposde")

	qui := "Felix Neto"

	fmt.Fprint(w, "A propos de...", qui)
}

// listerToutesCachacas é o endpoint para listar todas as cachaças cadastradas
func listerToutesCachacas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: listerToutesCachacas")
	fmt.Println(Cachacas)
	json.NewEncoder(w).Encode(Cachacas)
}

// listerUneCachaca é o endpoint para listar uma cachaça buscando pelo nome informado na URL
func listerUneCachaca(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: listerUneCachaca")

	vars := mux.Vars(r)
	cle := vars["nome"]
	cachacaencontra := false

	// fmt.Fprintf(w, "Nome selecionado: "+cle)

	for _, cachaca := range Cachacas {
		if strings.ToLower(cachaca.Nome) == strings.ToLower(cle) {
			cachacaencontra = true
			json.NewEncoder(w).Encode(cachaca)
		}
	}
	// Mensagem padrao para consulta vazia
	if cachacaencontra != true {
		fmt.Fprintf(w, "Nehuma cachaca encontrada com o nome: \""+cle+"\"")
	}
}

// debut1 é a função que vai ativar as rotas
func debut1() {

	roteur := mux.NewRouter().StrictSlash(true)

	roteur.HandleFunc("/", pageInitial)
	roteur.HandleFunc("/aproposde", aProposDe)
	roteur.HandleFunc("/listertoutescachacas", listerToutesCachacas)
	roteur.HandleFunc("/listerunecachaca/{nome}", listerUneCachaca)

	log.Fatal(http.ListenAndServe(":8085", roteur))
}

// Cachaca é a estrutura base para o ojeto cachaça
type Cachaca struct {
	Nome   string `json:"nome"`
	Volume string `json:"volume"`
	Custo  string `json:"custo"`
}

var Cachacas []Cachaca
