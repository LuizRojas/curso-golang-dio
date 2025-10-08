package main

import "fmt"

func main() {

	// x := 4

	// if x == 2 || x == 3 {
	// 	fmt.Println("Sim, x é 2 ou 3!")
	// } else {
	// 	fmt.Println("X não é 2 nem 3!")
	// }

	x := 6

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("X é múltiplo de 2 e 3")
	} else {
		fmt.Println("X não é múltiplo de 2 nem 3!")
	}
}
