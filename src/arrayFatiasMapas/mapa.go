package main

// mapas são coleções ordenadas (lista) pares chave-valor
func main() {
	// x := make(map[string]int)
	// x["chave"] = 10
	// fmt.Println(x["chave"])

	// x := make(map[int]int)
	// x[1] = 20
	// x[2] = 30
	// fmt.Println(x[1], x[2])

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Lítio"
	println(elemento["H"], elemento["He"], elemento["Li"])
}
