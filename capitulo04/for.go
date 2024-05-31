package main

import "fmt"

// func main() {
// 	i := 1 // variável chamada "i" armazenando o número que vamos imprimir.
// 	for i <= 10{ // "("i menor ou igual a 10"). loop for usando a palavra-chave for.
// 		fmt.Println(i)
// 		i = i + 1 // sem essa condição o nosso programa sempre seria avaliado como verdadeiro.
// 	}
// }

/*
Como exercício, vamos percorrer o programa como um computador faria:
1. Crie uma variável chamada i com o valor 1.
2. É <= 10? Sim.
3. Imprima i
4. Defina i como i + 1 (i agora é igual a 2).
5. É <= 10? Sim.
6. Imprima i.
7. Defina i como i + 1 (i agora é igual a 3).
8. ...
9. Defina i como i + 1 (i agora é igual a 11).
10. Sou <= 10? Não.
11. Não há mais nada a fazer, então saia.
*/

// O programa anterior poderia também ser escrito assim:

// func main() {
// 	for i := 1; i <= 10; i++{
// 		fmt.Println(i)
// 	}
// }


// Além de imprimir o prgrama vai dizer se é par ou ímpar

func main() {
	for i := 1; i<=10; i++{
		if i % 2 == 0 { // o resto de i / 2 é igual a 0?
			fmt.Println(i, "par")
        } else {
			fmt.Println(i, "ímpar")
		}
	}
}

/* Vamos percorrer este programa:
1. Crie uma variável i do tipo int e atribua a ela o valor 1.
2. É menor ou igual a 10? Sim: pule para o bloco if.
3. O resto de i ÷ 2 é igual a 0? Não: pule para o bloco else.
4. Imprima i seguido de ímpar.
5. Incremento i (a instrução após a condição).
6. É menor ou igual a 10? Sim: pule para o bloco if.
7. O resto de i ÷ 2 é igual a 0? Sim: pule para o bloco if.
8. Imprima i seguido de par e assim por diante até que i seja igual a 11.
9.…
*/