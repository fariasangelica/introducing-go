// 1 . O que o programa imprime?

// package main

// 	import "fmt"

// 	func main() {
// 		i :=10
// 		if i > 10 {
// 			fmt.Println("Grande")
// 		} else {
// 			fmt.Println("Pequeno")
// 		}

// 	}

// 2 . Escreva um programa que imprima todos os números entre 1 e 100 que sejam divisíveis por 3(ou seja, 3, 6, 9, etc.).

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

// Escreva um programa que imprima os números de 1 a 100, mas para múltiplos de 3 imprima "Fizz" em vez do número e para múltiplos de 5 imprima "Buzz". Paea números múltiplos de 3 e 5, imprima "FizzBuzz
