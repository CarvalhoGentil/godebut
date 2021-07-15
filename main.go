package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Iniciando alambique em http://localhost:8085 ")
	Cachacas = []Cachaca{
		Cachaca{Id: "0", Nome: "51", Volume: "974ml", Custo: "8"},
		Cachaca{Id: "1", Nome: "Matuta", Volume: "1000ml", Custo: "30"},
	}
	debut()
}

// pageInitial (Home Page) função para mostrar uma home page qualquer
func pageInitial(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: pageInitial")

	fmt.Fprintln(w, "- ......... -")
	fmt.Fprintln(w, "-- 2ez4Flz --")
	fmt.Fprintln(w, "-  GoDrink  -")
	fmt.Fprintln(w, "- ......... -")
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
	json.NewEncoder(w).Encode(Cachacas)
}

// listerUneCachaca é o endpoint para listar uma cachaça buscando pelo nome informado na URL
func listerUneCachaca(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: listerUneCachaca")

	vars := mux.Vars(r)
	cle := vars["nome"]
	cachacaencontra := false

	// Buscar na lista de cachacas o nome informado
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

// creerNouvelleCachaca é o endpoint para criar novos registros de cachaca
func creerNouvelleCachaca(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: creerNouvelleCachaca")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var cachaca Cachaca

	json.Unmarshal(reqBody, &cachaca)

	// Validar dados basicos antes de adicionar
	if cachaca.Id != "" && cachaca.Nome != "" {
		Cachacas = append(Cachacas, cachaca)

		json.NewEncoder(w).Encode(cachaca)

	} else {
		fmt.Fprintln(w, "Os dados de ID e NOME devem ser preenchidos !")
	}

}

// debut é a função que vai ativar as rotas
func debut() {

	roteur := mux.NewRouter().StrictSlash(true)

	// Rotas de visualizaçõ
	roteur.HandleFunc("/", pageInitial)
	roteur.HandleFunc("/aproposde", aProposDe)

	// Rotas de endpoits do CRUD
	roteur.HandleFunc("/listertoutescachacas", listerToutesCachacas)
	roteur.HandleFunc("/creernouvellecachaca", creerNouvelleCachaca).Methods("POST")
	roteur.HandleFunc("/listerunecachaca/{nome}", listerUneCachaca)

	log.Fatal(http.ListenAndServe(":8085", roteur))
}

// Cachaca é a estrutura base para o projeto cachaça
type Cachaca struct {
	Id     string `json:"id"`
	Nome   string `json:"nome"`
	Volume string `json:"volume"`
	Custo  string `json:"custo"`
}

var Cachacas []Cachaca
