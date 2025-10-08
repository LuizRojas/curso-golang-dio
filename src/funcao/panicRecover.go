// panic: erro do programador | erro execuão tempo
// recover: ela interrompe o panic

package main

import "fmt"

func main() {
	// panic("Pânico")
	// x := recover()
	// fmt.Println(x)

	defer func() {
		x := recover()  // terminar a função de imediato
		fmt.Println(x)
	}()
	panic("Pânico")
}
