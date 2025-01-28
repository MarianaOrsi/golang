package main

import "fmt"

// loops ou laços de repetição


// i++ -> soma 1 (i = i + 1)
// i-- -> subtraindo 1
// += -> é a mesma coisa que sum = sum + i 
// -= -> é a mesma coisa que sum = sum - i 

func main() {
	 sum := 0
	 for i := 0; i < 10; i++ {
		sum += 1
	 }
	 fmt.Println(sum)


	// for range
	 frutas := []string{"melancia", "melão", "morango", "uva", "maçã"}
	 for _, fruta := range frutas {
		fmt.Println("Fruta:", fruta)
	 }
}