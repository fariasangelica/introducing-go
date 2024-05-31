// Programa que converta pés em metros (1 pé = 0,3048 m).

package main

import "fmt"

func main() {
	fmt.Println("Digite um número:")
	var input float64
	fmt.Scan("%f", &input)

	output := input * 0.3048

	fmt.Println(output)
}
