// package main

// import "fmt"

// func main() {
// 	const x string = "Hello, World"
// 	fmt.Println(x)
// }


// Definindo Múltiplas Variáveis
// var(
// 	a = 5
// 	b = 10
// 	c = 15
// ) //Use a palavra-chave var (ou const) seguida de parênteses com cada variável em sua própria linha.


// Exemplo de programa
package main

import "fmt"

func main() {
	fmt.Println("Digite um número:") 
	var input float64
	fmt.Scanf("%f, &input") // Função do pacote "fmt". Basicamente, ele preenche a entrada com o número que inserimos.
	
	output := input * 2

	fmt.Println(output)
}


