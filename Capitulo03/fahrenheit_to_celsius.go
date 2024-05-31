// Programa que converta de Fahrenheit para Celsius (C = (F y 32) * 5/9).

package main

import "fmt"

func main() {
	fmt.Println("Digite um número:")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5/9
	fmt.Println(output)
}