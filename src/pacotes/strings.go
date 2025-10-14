// caixinha de funções
// função contains
// contains: procura dentro de uma string outra string menor
// ex: Terra - rr (true)
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("computador", "comp"))
	fmt.Println(strings.Count("maquina", "a"))  // conta quantas vezes 'a' aparece em 'maquina'
}
