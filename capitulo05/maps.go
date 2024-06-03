// var x map[string]int // x é um map de string para ints

// O tipo de map é representado pela palavra-chave map, em seguida pelo tipo de chave entre colchetes e por fim pelo tipo de valor



// Podemos excluir itens de um map usando a função de exclusão:

// excluir (x, 1)

package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
}


// Uma maneira mais curta de criar mapas:

 elements := map[string]string{
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

}


 // Programa de elementos quimicos com nome, estado padrão (estado em temperatura ambiente)

 package main

 import "fmt"

 func main() {
	elements := map[strings]map[string] string{
		"H": map[string]string{
			"name":"Hydrogen"
			"state" : "gas"
		},
		"He": map[string]string{
			"name:""Helium"
			"state":"solid",
		},
		"Li": map[string]string{
			"name":"Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":"Beryllium",
			"state":"solid",
		},
		"B": map[string]string{
			"name":"Boron",
		    "state":"solid",
		},
		"C": map[string]string{
			"name":"Carbon"
			"state":"solid"
		},
		"N": map[string]string{
			"name":"Nitrogen",
			"state":"gas",
		},
		"O": map[string]string{
			"name":"Oxygen",
			"state":"gas"
		},
		"F": map[string]string{
			"name":"Fluorine"
			"state":"gas"
		},
		"Ne": map[string]string{
			"name":"Neon",
			"state":"gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
	
	    // Ainda veremos uma maneira melhor de armazenar informações estruturadas.

 }
