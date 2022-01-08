package main()

import (
	"os"
	"fmt"
)

func main() {
	// PRINTAR UMA VARIAVEL DE AMBIENTE
    fmt.Println("VAR NAME : ", os.Getenv("GOPATH"))

	// PRINTAR TODAS AS VARIAVEIS DE AMBIENTE DO SISTEMA
	for _, env := range os.Environ(){
        fmt.Println(env)
    }
}
