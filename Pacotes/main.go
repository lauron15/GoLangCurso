package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Olá, Mundo!1")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("lauroafp@gmail.com")
	// verificação simples, se o erro for diferente de nil (null do go)
	if erro != nil {
		fmt.Printf("Erro: %v\n", erro)
	} else {
		fmt.Println("O email está em um formato válido!")
	}

}

//Eu fiz a mesma ideia que o ReactJs então, eu criei uma função,
// em um objeto diferente, e chamei essa funçãp no principal.
// primeiro no import, com o modulo/nome do arquivo.
//e depois na função com nome do arquivo.Nome da função()
