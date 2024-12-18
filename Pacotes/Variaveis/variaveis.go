package main

import "fmt"

func main() {

	var variavel string = "variávevel1"
	variavel2 := "variável2"
	fmt.Println(variavel)
	fmt.Println(variavel2)

}

// variáveis, podem ou não ter o tipo explicito. ou seja, posso alocar o tipo diretamente, como por exemplo: var variavel string ="...",
//ou, variavel2:= "variável2" atribuir diretamente, sem especificar o tipo, contudo, o próprio Go, encontra o tipo para mim.
