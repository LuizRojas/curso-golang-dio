// Escreva um código em Go utilizando todo o conhecimento adquirido até o momento! E nesse código você precisará,
// baseado em nossas aulas de concorrência, utilizar canais e goroutines para que o seu programa exiba, em alternância,
// as palavras ping e pong.

package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func imprimir(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go ping(c)
	go pong(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
