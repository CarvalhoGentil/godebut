package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"environ"
)

func main() {
	fmt.Println("Iniciando alambique em http://localhost:8085 ")
	Cachacas = []Cachaca{
		{Id: "0", Nome: "51", Volume: "974ml", Custo: "8"},
		{Id: "1", Nome: "Matuta", Volume: "1000ml", Custo: "30"},
	}

	initialeMigration()

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

	db, err := gorm.Open("postgres", "host=teste user=teste password=teste dbname=teste port=5432 sslmode=disable")
	if err != nil {
		fmt.Println("Erro de conexão ao banco de dados: \"" + err.Error() + "\"")
	}
	defer db.Close()

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

// initialeMigration é a função para executar a migração no banco
func initialeMigration() {

	db_try := 40
	fmt.Println("Capturando variaveis de ambiente para conectar ao banco de dados...")

	db_host := environ.GetEnvironValue("DEV_DB_HOST")
	db_user := environ.GetEnvironValue("DEV_DB_USER")
	db_passwd := environ.GetEnvironValue("DEV_DB_PASSWD")
	db_name := environ.GetEnvironValue("DEV_DB_NAME")
	db_port := environ.GetEnvironValue("DEV_DB_PORT")

	var db_url = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		db_host,
		db_user,
		db_passwd,
		db_name,
		db_port)

	fmt.Println("Conectando na base", db_name)

	// db_url := ("host=db user=teste password=teste dbname=teste port=5432 sslmode=disable")

	db := PostgresConn(db_url, db_try)
	fmt.Println("Conexão OK")

	defer db.Close()

	db.AutoMigrate(&Consumidor{})

	// Inserção de dados para teste
	db.Create(&Consumidor{Nome: "felix marmotinha", Idade: "18"})
	db.Create(&Consumidor{Nome: "felix-2-devops", Idade: "33"})
	db.Create(&Consumidor{Nome: "jorbson-2-scripts", Idade: "21"})
	db.Create(&Consumidor{Nome: "paulo-2-manager", Idade: "28"})
	db.Create(&Consumidor{Nome: "Nicer", Idade: "1"})
	db.Create(&Consumidor{Nome: "Ezy", Idade: "2"})
	db.Delete(&Consumidor{}, 1)

}

// PostgresConn tenta conectar e retornar uma conexão ao banco de dados postgres com retry
// Se o numero de retry esgotar, gera panic(err)
func PostgresConn(connectionUrl string, retry int) *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open("postgres", connectionUrl)
	for err != nil {
		fmt.Println("Tentativa conexão ao banco...", retry)

		if retry > 1 {
			retry--
			time.Sleep(5 * time.Second)
			db, err = gorm.Open("postgres", connectionUrl)
			continue
		}
		panic(err)
	}
	return db
}

// Environ deve printar variaveis de ambiente do sistema
func Environ(var_name string) {
	// PRINTAR UMA VARIAVEL DE AMBIENTE
	fmt.Println("Printando todas as variaveis do ambiente")

	// PRINTAR TODAS AS VARIAVEIS DE AMBIENTE DO SISTEMA
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}

// GetEnvironValue busca o valor de uma variavel de ambiente
func GetEnvironValue(var_name string) string {

	// RETORNAR VALOR DA VARIAVEL CONSULTADA
	var env_result string
	env_result = os.Getenv(var_name)
	return env_result
}

// debut é a função que vai ativar as rotas da API
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

type Consumidor struct {
	gorm.Model
	Nome  string
	Idade string
}

var Cachacas []Cachaca
