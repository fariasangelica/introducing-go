var x []float64

// x foi criado com o comprimento zero.
// se quiser criar um slices, deve usar a função interna make

x := make([]float64, 5, 10)

// os slices são sempre associadas a um array e, embora nunca possam ser maiores que o array, podem ser menores.
// função make também permite um terceiro parâmetro

x := make([]float64, 5, 10) // 10 representa a capacidade do array.


// outra forma de criar slices é usando a expressão [low : high]:

arr := [5]float64{1,2,3,4,5}
x := arr[0:5]

// low é o índice de onde iniciar a fatia.
// arr[0:] é o mesmo que arr[0:len(arr)]
// arr[:5] é o mesmo que arr[0,5]
// arr[:] é o mesmo que arr[0:len(arr)] 

func main() {
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5) // apprend cria um novo slice pegando uma fatia existente e anexando todos os argumentos seguintes.
	fmt.Println(slice1, slice2)     // slice1 possui [1,2,3] e slice2 possui [1,2,3,4,5]
}


// Exemplo de Copy

func main() {
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2) // slice1 possui[1,2,3] e slice2 possui [1,2]
}   // O conteúdo de slice1 é copiado para slice2, mas como slice2 tem espaço apenas para dois elementos, apenas os dois primeiros elementos de slice1 são copiados.
