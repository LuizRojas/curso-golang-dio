package main

import "fmt"

// Exibir todos os números entre 1 e 100 que são divisíveis por 3
func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
