package auxiliar

import "fmt"

//Como boa pratico no Go, sempre que uma função for exportada
//é bom deixar um comentário em cima, explicando o que ela faz.

func Escrever() {
	fmt.Println("Escrevendo do arquivo pacote auxiliar")
	escrever2()
}
