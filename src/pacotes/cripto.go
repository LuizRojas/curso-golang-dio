package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	// criar a hash
	h := sha1.New()
	
	// escrever dados para o hash
	h.Write([]byte("Código do pacote cripto"))

	// calcular o hash
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
