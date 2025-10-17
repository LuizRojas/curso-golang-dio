package main

import (
	"log"
	"os"
)

func main() {
	message := []byte("Hello, Gophers!")
	err := os.WriteFile("Teste.txt", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
