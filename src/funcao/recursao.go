package main

import "fmt"

func fatorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fatorial(n-1)
	}
}

func main() {
	valor := 5
	fmt.Printf("O fatorial de %d é: %d\n", valor, fatorial(valor))
}
