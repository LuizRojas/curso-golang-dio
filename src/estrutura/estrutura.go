package main

import "fmt"

// são coleções de "campos"
// agrupar dados
// formar registros
// criar um novo tipo de dado

type pessoa struct {
	nome   string
	idade  int
	altura float64
}

func main() {
	fmt.Println(pessoa{"Ana", 54, 1.61})
	fmt.Println(pessoa{"João", 34, 1.76})
}
