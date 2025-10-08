package main

import "fmt"

// função também é chamada de procedimento ou sub-rotina
// parte do código

func media(lista []float64) float64 {
	total := 0.0

	for _, valor := range lista {
		total += valor
	}

	return total / float64(len(lista))
}

func main() {
	lista := []float64{98, 76, 88, 81, 59}
	fmt.Println(media(lista))
}
