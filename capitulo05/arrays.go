package main

import "fmt"

// func main() {
// 	var x [5]int
// 	x[4] = 100
// 	fmt.Println(x)
// }

// x[4] = 100 deve ser lido "defina o quinto elemento da matriz x como 100"
// x[4] representa o quinto elemento em vez de quarto, assim como as strings, os arrays são indexados a partir de 0.
// poderia ser mudado o fmt.Println(x) para fmt.Println(x[4]).



// Este programa calcula a média de uma série de pontuações de testes.
// func main() {
// 	var x [5]float64
// 	x[0] = 98
// 	x[1] = 93
// 	x[2] = 77
// 	x[3] = 82
// 	x[4] = 83

// 	var total float64 = 0
// 	for i := 0; i < 5; i++ {
// 		total += x[i]
// 	}
// 	fmt.Print(total / 5)
// }

// Primeiro, criamos um array de comprimento 5 para armazenar nossas pontuações de teste e, em seguida, preenchemos cada elemento com uma nota.
// Configuramos um loop para calcular a pontuação total.
// Por fim, dividimos a pontuação total pela quantidade de elementos para encontrar a média.

// Melhorando esse código
// func main() {
// 	var x [5]float64
// 	x[0] = 98
// 	x[1] = 93
// 	x[2] = 77
// 	x[3] = 82
// 	x[4] = 83

// 	var total float64 = 0
// 	for i := 0; i < len(x); i++ {
// 		total += x[i]
// 	}
// 	fmt.Print(total / float64(len(x))) // Exemplo de conversão de tipo. Usa-se o nome do tipo como uma função.
// }


// Outra forma

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for _, value := range x { // Sublinhado (_) é usado para informar ao compilador que não precisamos disso (não precisamos da variável iteradora).
		total += value
	}
	fmt.Print(total / float64(len(x))) // Exemplo de conversão de tipo. Usa-se o nome do tipo como uma função.
}

// Sintaxe mais curta para a criação da variável

x := [5]float64{98, 93, 77, 82, 83} // para remover um elemento do array é só comentar: 82, //83...

// Adicionar ou remover elementos do arrays também requer que alteremos o comprimento dentro dos colchetes. Limitações e por isso não se ver arrays sendo usado nos códigos.
// vemos mais slices.

