package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%2 == 0 && i%3 == 0:
			fmt.Println("Pin Pan")
		case i%2 == 0:
			fmt.Println("Pan")
		default:
			fmt.Println("Pin")
		}
	}
}
