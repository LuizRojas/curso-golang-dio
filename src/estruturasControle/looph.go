package main

import "fmt"

func main() {

	for horas := 0; horas < 12; horas++ {
		for minutos := 0; minutos < 60; minutos++ {
			for segundos := 0; segundos <= 60; segundos++ {
				fmt.Printf("Hora: %d:%d:%d\n", horas, minutos, segundos)
			}
		}
	}
}
