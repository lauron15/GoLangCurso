package auxiliar

import "fmt"

func escrever2() {
	fmt.Printf("Deu certo")
}

//dentro do proprio pacote eu consigo referenciar as funções que estão com letra minúscula.
// no exemplo, eu chamei a função do auxiliar2 no auxiliar 1.
//sendo assim, na tela, eu vejo tanto o que foi impresso na auxilair, quanto na auxiliar 2
