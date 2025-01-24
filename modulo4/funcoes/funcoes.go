package main

import "fmt"

func main() {
	fmt.Println(soma(42, 13))

	// somaDosValores := soma(42, 13)
	// fmt.Println(somaDosValores)

	sub := subtracao(10, 5)
	fmt.Println(sub)
}

func soma(x, y int) int {
	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}

/*
Função que começa com a letra minúscula é PRIVADA
Ela só pode ser usada dentro do próprio pacote
Ex: func printaNomeCompleto
*/


/*
Função que começa com a letra maiuscula é PÚBLICA
Ela pode ser usada fora do próprio pacote
Ex: main.PrintaNomeCompleto
*/
