package main

import (
	"fmt"
)

func main() {

	// bool (true/false)
	fmt.Printf("Type: %T - Value: %v\n", true, true)

	// string - sequência de bytes
	fmt.Printf("Type: %T - Value: %v\n", "mari", "mari")
	fmt.Printf("Type: %T - Value: %v\n", "1", "1")

	// int
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)

	//float (float64 / float32) - decimal
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)

}

/* Tipos:

- bool (true / false)
- string - sequência de bytes
- int
- float (float64 / float32) - decimal

%v o valor em um formato padrão
%T uma representação de sintaxe Go do tipo do valor

*/
