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

// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i % 3 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }


// 3. Escreva um programa que imprima os números de 1 a 100, mas para múltiplos de três imprima "Fizz" em vez do número 
// e para múltiplos de cinco imprima "Buzz". Para números múltiplos de três e cinco, imprima "FizzBuzz".

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		 } else if i%3 == 0 {
			fmt.Println("Fizz")
		 } else if i%5 == 0 {
			fmt.Println("Buzz")
		 } else {
			fmt.Println(1)
		 }
	}
}	
		 

		
	





