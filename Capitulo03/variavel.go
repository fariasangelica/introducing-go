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

var x string = "hello" var y
string = "world"
fmt.Println(x==y) // imprimirá falso, pois hello não pe o mesmo que wolrd.

var x string = "hello" var y
string = "hello"
fmt.Println(x==y) //imprime verdadeiro porque as duas strings são iguais.