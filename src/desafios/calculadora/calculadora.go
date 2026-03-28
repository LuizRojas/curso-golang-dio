package main

import "fmt"

func soma(e ...int) int {
	total := 0
	for _, v := range e {
		total += v
	}
	return total
}

func subtracao(e ...int) int {
	total := 0
	for _, v := range e {
		total -= v
	}
	return total
}

func multiplicacao(e ...int) int {
	total := 1
	for _, v := range e {
		total *= v
	}
	return total
}

func divisao(e ...float64) float64 {
	var total float64 = 1.0
	for _, v := range e {
		total = v / total
	}
	return total
}

func main() {
	a := soma(6, 7)
	b := subtracao(171, 157)
	c := multiplicacao(3, 2, 1)
	d := divisao(20, 40)
	fmt.Println(a, b, c, d)
}