package main

import "fmt"

type retangulo struct {
	comprimento, altura float64
}

func (r *retangulo) area() float64 {
	return r.comprimento * r.altura
}

func (r *retangulo) perimetro() float64 {
	return 2*r.comprimento + 2*r.altura
}

func main() {
	r := retangulo{comprimento: 23.7, altura: 12.5}

	fmt.Println("Área do retângulo:", r.area())
	fmt.Println("Perímetro do retângulo:", r.perimetro())
}
