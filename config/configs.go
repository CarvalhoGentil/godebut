package main

import (
	"fmt"
	"path/filepath"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Dbuser string
	Dbname string
	Dbpass string
}

type MyInterface interface{}

func main() {
	fmt.Println("Definindo configurações")

	start_config()
}

func return_struct(user string, name string, pass string) MyInterface {
	// Validar esse retorno como no exemplo
	// https://play.golang.org/p/tHbTZHEQZ-L
	return &Configuration{
		Dbuser: user,
		Dbname: name,
		Dbpass: pass,
	}
}

func start_config() {

	var my_config_path string
	my_config_path, _ = filepath.Abs("config.prod.json")
	fmt.Println(my_config_path)

	var my_config Configuration
	err := gonfig.GetConf(my_config_path, &my_config)
	// err := gonfig.GetConf("/home/felix/cofe/4fun/godebut/config/config.prod.json", &my_config)
	if err != nil {
		panic(err)
	}
	my_config{Dbuser: my_config.Dbuser, my_config.Dbname, my_config.Dbpass}
	fmt.Println(my_config.Dbpass, my_config.Dbpass, my_config.Dbpass)
	fmt.Println(my_config)
	// return my_config.Dbpass

}
