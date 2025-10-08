package main

import "fmt"

func main() {
	// var x []float64
	// fatia:= make([]float64, 4)
	// fatia := [low:high]
	// fatia:= arr[0:5]
	// acrescentar elemento no final da fatia

	// arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	// fatia := arr[0:7] // make([]float64, 5)
	// fmt.Println(fatia)

	fatia1 := []int{1, 2, 3}
	fatia2 := make([]int, 1)
	copy(fatia2, fatia1)
	fmt.Println(fatia1, fatia2)
}
