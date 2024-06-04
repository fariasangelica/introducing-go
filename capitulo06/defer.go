// Este programa imprime o primeiro seguido do segundo. Basicamente, defer move a chamada para o segundo no final da função

package main 

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second()
	first()
}

// Defer é frequentemente usado quando os recursos precisam ser liberados de alguma forma. Por exemplo, abrimos um arquivo, precisamos fechá-lo mais tarde.

f, _ := os.Open(filaname)
defer f.close()

// Mantém o nosso Close call perto do nosso Open call para que seja mais fácil de entender.
// Se a nossa função tivesse múltiplas instruções de retorno.
// As funções adiadas são executadas mesmo se ocorrer um panic nop tempo de execução.

