package main

import "fmt"

const ebulicaoF float64 = 212.0

func main() {
	// var tempF float64 = ebulicaoF
	// var tempC float64 = (tempF - 32) * 5 / 9

	// fmt.Println("A temperatura de ebulição da água em °F = ", tempF)
	// fmt.Println("A temperatura de ebulição da água em °C = ", tempC)

	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9

	fmt.Printf("A temperatura em °F = %.2f\nA temperatura em °C = %.2f", tempF, tempC)
}
