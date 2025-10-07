package main

import "fmt"

const ebulicaoKelvin float64 = 373

func main() {
	tempK := ebulicaoKelvin
	tempC := tempK - 273

	fmt.Printf("A água entra em ebulição em %.2f°K\nJá em Celsius: %.2f°C", tempK, tempC)

}
