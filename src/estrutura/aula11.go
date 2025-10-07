package main

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int8
}

func mostrarPessoa(Pessoa p) {
	println("Nome:", p.Nome)
	println("Sobrenome:", p.Sobrenome)
	println("Idade:", p.Idade)
}

func main() {
	pessoa1 := Pessoa{"Luiz", "Rojas", 20}
	pessoa2 := Pessoa{"Ana", "Castella", 22}
	mostrarPessoa(pessoa1)
	println()
	mostrarPessoa(pessoa2)
}
