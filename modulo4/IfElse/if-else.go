package main

import "fmt"

func main() {

	valor := 1
	if valor == 1 {
		fmt.Println("Valor é igual a 1")
	} else {
		fmt.Println("Valor não é igual a 1")
	}

	numero := 9
	if numero != 7 {
		fmt.Println("O numero é diferente de 7")
	} else {
		fmt.Println("O numero é 7")
	}

	// if - else if - else

	outroNumero := 2
	if outroNumero == 2 {
		fmt.Println("O outro numero é igual a 2")
	} else if outroNumero == 1 {
		fmt.Println("O outro numero é igual a 1")
	} else {
		fmt.Println("O outro numero é diferente de 1 e 2")
	}

	restoNum := 7
	if restoNum%2 == 0 {
		fmt.Printf("%d é par", restoNum)
	} else {
		fmt.Printf("%d é impar", restoNum)
	}


/* 

### Operações
Soma: 1 + 1
Subtração 2 - 1
Divisão: 10 / 2
Multiplicação: 2 * 2
Resto da divisão por x: 7%2 (resto da divisao por 2)

### Operadores de comparação
> (maior): Retorna verdadeiro caso o primeiro valor seja maior que o segundo.
>= (maior ou igual): Retorna verdadeiro caso o primeiro valor seja maior ou igual ao segundo.
< (menor): Retorna verdadeiro caso o primeiro valor seja menor que o segundo.
<= (menor ou igual): Retorna verdadeiro caso o primeiro valor seja menor ou igual ao segundo
== (igual a): Retorna verdadeiro caso o primeiro valor seja igual ao segundo.
!= (diferente de): Retorna verdadeiro caso o primeiro valor seja diferente do segundo.

*/
	
}