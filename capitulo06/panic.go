// package main 

// import "fmt"

// func main() {
// 	panic("PANIC")
// 	str := recover() // this will never happen
// 	fmt.Println(str)
// }

// Podemos lidar com um pânico em tempo de execução com a função de recuperação integrada.
// Precisamos emparelhá-lo com defer porque sem a chamada para panic interrompe imediatamente a execução da função.
package main 

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

}