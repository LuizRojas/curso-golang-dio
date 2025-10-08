package main

import "fmt"

func main() {
	// imprimir apenas números pares
	
	for i := 0; i <= 20; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
}
