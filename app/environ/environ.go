package environ

import (
	"fmt"
	"os"
)

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
