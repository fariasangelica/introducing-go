// Como acessar o quarto elemento de um array ou slice?
// arr [3]

// Qual é o comprimento de uma fatia criada usando make([]int, 3, 9)?
// 3

// Dado o seguinte array, o que x[2:5] lhe daria?
     // x := [6]string{"a","b","c","d","e","f"}
// [c d e]

// Escreva um programa que encontre o menor número nesta lista:
    // x := []int{
	// 	48,96,86,68,
	// 	57,82,63,70,
	// 	37,34,83,27,
	// 	19,97,9,17,
	//  }

package main 

import "fmt"

func main() {
	var min int
	x := []int{
		48,96,86,68,
	 	57,82,63,70,
	 	37,34,83,27,
	 	19,97,9,17,
	}
	for i, v := range x {
		if i == 0 || v < min {
			min = v
		}
	}
	fmt.Println(min)
}