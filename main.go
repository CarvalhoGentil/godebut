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
		{Id: "0", Nome: "51", Volume: "974ml", Custo: "8"},
		{Id: "1", Nome: "Matuta", Volume: "1000ml", Custo: "30"},
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

// toutesCachacas é o endpoint para listar todas as cachaças cadastradas
func toutesCachacas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: toutesCachacas")
	json.NewEncoder(w).Encode(Cachacas)
}

// uneCachaca é o endpoint para listar uma cachaça buscando pelo nome informado na URL
func uneCachaca(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: uneCachaca")

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
		fmt.Fprintln(w, "Nenhuma cachaca encontrada com o nome: \""+cle+"\"")
	}
}

// nouvelleCachaca é o endpoint para criar novos registros de cachaca
func nouvelleCachaca(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: nouvelleCachaca")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var cachaca Cachaca

	json.Unmarshal(reqBody, &cachaca)

	// Validar dados basicos antes de adicionar
	if cachaca.Id != "" && cachaca.Nome != "" {
		Cachacas = append(Cachacas, cachaca)
		fmt.Fprintln(w, "A cachaca \""+cachaca.Nome+"\" foi adicionada a lista.")

		json.NewEncoder(w).Encode(cachaca)

	} else {
		fmt.Fprintln(w, "Os dados de ID e NOME devem ser preenchidos !")
	}
}

// renouvelleCachaca ó o endpoint para aualizar um registro de cachaca
func renouvelleCachaca(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: renouvelerCachaca")

	vars := mux.Vars(r)
	cle := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var update_cachaca Cachaca

	json.Unmarshal(reqBody, &update_cachaca)

	cachacaencontra := false

	if cle != "" {
		// Buscar na lista de cachacas o Id informado
		for index, cachaca := range Cachacas {
			if strings.ToLower(cachaca.Id) == strings.ToLower(cle) {
				cachacaencontra = true

				// Atualizar informações da cachaca na lista
				Cachacas[index] = update_cachaca

				fmt.Fprintln(w, "A cachaca com Id \""+cle+"\" foi aualizada.")
			}
		}
		// Mensagem padrão para consula vazia
		if cachacaencontra != true {
			fmt.Fprintln(w, "Nenhuma cachaca encontrada com o Id: \""+cle+"\"")
		}
	} else {
		fmt.Fprintln(w, "Os dados de ID deve ser preenchido !")
	}
}

// effacerCachaca é o endpoint para deletar
func effacerCachaca(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: effacerCachaca")

	vars := mux.Vars(r)
	cle := vars["nome"]
	cachacaencontra := false

	// Buscar na lista de cachacas o nome informado
	for index, cachaca := range Cachacas {
		if strings.ToLower(cachaca.Nome) == strings.ToLower(cle) {
			cachacaencontra = true

			// Remover a cachaca informada da lista
			Cachacas = append(Cachacas[:index], Cachacas[index+1:]...)

			fmt.Fprintln(w, "A cachaca \""+cle+"\" foi removida da lista.")
		}
	}
	// Mensagem padrão para consula vazia
	if cachacaencontra != true {
		fmt.Fprintln(w, "Nenhuma cachaca encontrada com o nome: \""+cle+"\"")
	}
}

// debut é a função que vai ativar as rotas
func debut() {

	roteur := mux.NewRouter().StrictSlash(true)

	// Rotas de visualização
	roteur.HandleFunc("/", pageInitial)
	roteur.HandleFunc("/aproposde", aProposDe)

	// Rotas de endpoints do CRUD
	roteur.HandleFunc("/v1/toutescachacas", toutesCachacas)
	roteur.HandleFunc("/v1/unecachaca", nouvelleCachaca).Methods("POST")
	roteur.HandleFunc("/v1/unecachaca/{id}", renouvelleCachaca).Methods("PUT")
	roteur.HandleFunc("/v1/unecachaca/{nome}", effacerCachaca).Methods("DELETE")
	roteur.HandleFunc("/v1/unecachaca/{nome}", uneCachaca)

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
