package main

import (
	"fmt"
	"hash/crc32"  // "ola bom dia"  -> +/- "13245453461243123" (ou seja, transforma em números)
)

func main() {
	// criar a hash
	h := crc32.NewIEEE()

	// escrever dados para o hash
	h.Write([]byte("Código com pacote hash"))

	// calcular o hash
	v := h.Sum32()
	fmt.Println(v)
}
