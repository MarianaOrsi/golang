package main

import "fmt"

func main() {
	// Array - Tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	numPrimos := [7]int{2, 3, 5, 7, 11, 13, 29}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:3]) // os dois pontos ":" pega antes da posição 3, sem inluir o 3
	// ou seja, pega a posição 0, 1 e 2

	// Slices
	//var  slice []string
	slice := make([]string, 7) // criando slice do zero
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice)
	fmt.Println(len(slice))

	nunPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(nunPares)

	nunPares =append(nunPares, 14) // adicionando o número 14
	fmt.Println(nunPares)
}







 /* LISTAS

 1 - Arrays e Slices: Homogêneos
 todos os elementos tem o mesmo tipo
 [1, 2, 3, 4, 5, 6, 7] - []int
 ["mari", "orsi", "golang"] - []string

 2 - Maps: Heterogêneos
 pode misturar tipos
 estrutura chave - valor
 [key] = value
 chave tem um tipo, e o valor pode ter outro
 map[string]int
 { "mari": 29, "salem": 7}
 map[string]string
 { "mari": "orsi", "salem": "gato" }


 Array

 Tamanho fixo, de zero ou mais elementos do mesmo tipo
 Acessamos os valores com índice: a[0], a[1]...
 Função embutida len() retorna o tamanho do array
 Por conta do tamanho fixo, não é tão usado. Só em casos específicos

 Slice

 Tipo o array, mas sem tamanho fixo
 Acessamos os valores com índice: a[0], a[1]...
 Função embutida len() retorna o tamanho do slice
 Função append() usada para adicionar valores nesse slice */

 // https://ifrnead.github.io/rubynapratica/assets/images/estrutura-array.png