package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

// x[4] = 100 deve ser lido "defina o quinto elemento da matriz x como 100"
// x[4] representa o quinto elemento em vez de quarto, assim como as strings, os arrays são indexados a partir de 0.
// Poderia ser mudado o fmt.Println(x) para fmt.Println(x[4]).

