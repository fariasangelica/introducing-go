package main

import 

// As funções começam com a palavra-chave func, seguida do nome da função.
// os parâmetros (entradas) da função s]ao definidos: tipo de nome, tipo de nome...
// A função tem um parâmetro (a lista de pontuações) que chamamos de xs. Após os parâmetros
// colocamos o tipo retorno. Parâmentros e retornos são conhecidos como assiantura da função.

func average(xs []float64) float64 {

	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
} 

//panic("Not Implemented")
	//xs := []float{98,93,77,82,83}


