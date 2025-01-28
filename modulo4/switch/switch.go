package main

import (
	"fmt"
)


func main() {
	posicao := 2
	switch posicao {
	case 1: 
		fmt.Println("Primeiro lugar")
	case 2:
		fmt.Println("Segundo lugar")
	case 3: 
		fmt.Println("Terceiro lugar")
	}

	nome := "Homem-Aranha"
	switch nome {
	// nome == "Mariana"
	case "Mariana":
		fmt.Println("É uma pessoa")
	case "Salem":
		fmt.Println("É um gato")
	case "Homem-Aranha":
		fmt.Println("É um personagem")
	}
}