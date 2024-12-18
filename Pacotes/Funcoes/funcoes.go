package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	// aqui no (int8,int8) tenho dois, pois abaixo, vou retornar soma e subtração, logo, eu tenho que especificar um tipo para cada retorno, por isso os dois.
	// A abertura da função tem que estar na mesma linha da função por si só e não abaixo.

	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

// percebe-sse que o último int8, não é uma variável, e sim, a declaração do tipo de retorno a devolver.
// pois: No Go, se uma função retorna um valor, é obrigatório declarar o tipo de retorno no cabeçalho da função.
func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("texto da função 1")
	fmt.Println(resultado)

	// Então, eu posso retornar dois resultados a partir do mesmo parâmetro.
	//posso declarar duas váriáveis, e atribuir o valor as duas.
	//e quando der o print, chamar as duas.
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 50)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
