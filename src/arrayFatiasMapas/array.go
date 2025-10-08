package main

import "fmt"

func main() {
	var x [5]float64 // tipos complexos diferente de int, float, bool que são simples

	x[0] = 4.2
	x[1] = 5.1
	x[2] = 7.0
	x[3] = 2.8
	x[4] = 4.9

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	media := total / float64(len(x))

	fmt.Printf("A média das notas é: %.2f\n", media)
}
