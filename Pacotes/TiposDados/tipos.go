package main

import (
	"errors"
	"fmt"
)

func main() {

	//inteiros: existem 4 tipos de inteiros em Go. 8 / 16 / 32/ 64 (quantidade de bytes, igual o java)
	// o int também aceita número negativos
	// se eu não declarar a quantidade de bits, o computador aloca por si mesmo.

	//numero inteiros
	var numero int = 10000
	fmt.Println(numero)

	// uint, é o tipo que não aceita número negativos

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// numero reais
	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	numeroReal3 := 123456.7
	fmt.Println(numeroReal3)
	// Strings
	var str string = "Texto"
	fmt.Println(str)

	char := 'L'
	fmt.Println(char)
	//vai me devolver a posição e não a string.

	var erro error
	fmt.Println(erro) //devolve um nil (nulo)

	//como fazer o erro devolver uma mensagem?? código abaixo.
	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)

}
