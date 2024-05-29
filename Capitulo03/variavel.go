package main

import "fmt"

func main() {
	var
	nome string = "angelica farias" 
	fmt.Println(nome)
}

// Variáveis em Go são criadas primeiro usando a palavra-chave var , depois especificando o nome da variável (nome) 
// e o tipo (string) e, finalmente, atribuindo um valor à variável (angelica).

// Atribuir um valor à variável é opcional

// package main

// import "fmt"

// func main() {
// 	x string x = "Hello, World" // "x" recebe a string Hello, World
// 	fmt.println(x)
// }

// Usamos símbolo == é igualdade.

// var x string = "hello" 
// var y string = "world"
// fmt.Println(x==y) // imprimirá falso, pois hello não pe o mesmo que wolrd.

// var x string = "hello" var y
// string = "hello"
// fmt.Println(x==y) //imprime verdadeiro porque as duas strings são iguais.

// x:="Hello, World" 
// compilador Go é capaz deinferir o tipo com base no valor que atribuimos à variável.

// var x= "Hello, World" // inferência com a instrução var

// x.=5
// fmt.Println(x)


// Nomeando uma variável

// nome := "Max"
// fmt.Println("O nome do meu cachorro é:", nome)


// dogName := "Max"
// fmt.println("O nome do meu cachorro é", dogName)


// Escopo

package main

import "fmt"

func main() {
	var x string = "Hello, World"
	fmt.Println(x)
}

// Outra forma de escrever este programa seria assim:

package main

import "fmt"

var x string = "Hello, World"

func main() {
	fmt.Println(x)  // movi a variável para fora da função principal
}                   // isso significa que outras funções podem acessar está variável
    




