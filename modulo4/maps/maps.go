package main

import "fmt"


// formas de criar o map

// forma 1
func main(){
	idade := map[string] int{}
	idade["Mari"] = 29
	idade["Salem"] = 2
	fmt.Println(idade)
	fmt.Println(idade["Mari"])
	fmt.Println(idade["Salem"])


// forma 2	
	anoNasc := map[string] int {
		"Mari": 1995,
		"Salem": 2022,
	}

	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Mari"])
	fmt.Println(anoNasc["Salem"])

	// adicionando elemento ao map
	anoNasc["Eliana"] = 1975
	fmt.Println(anoNasc)
}

/*
--- Maps: Heterogêneos
 pode misturar tipos
 estrutura chave - valor
 [key] = value
 chave tem um tipo, e o valor pode ter outro
 map[k]v -> k = chave, v = valor

 map[string]int
 { "Mari": 29, "bento": 2}
 map[string]string
 { "Mari": "Orsi", "Salem": "Orsi" }
*/