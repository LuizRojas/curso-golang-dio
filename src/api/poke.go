package main

// GET: listagem de registros
// POST: adiciona um novo registro
// DELETE: remover um registro
// PUT e PATCH: editar
// Response
// Pokemon
//

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	nome string "json: 'name'"
	pokemon []Pokemon "json: 'pokemon_entries'"
}

type Pokemon struct {
	numero int "json: 'entry_number'"
	especie PokemonSpecies "json: 'pokemon_species'"
}

type PokemonSpecies struct {
	nome string "json: name"
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/") // mapeia

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.nome)
	fmt.Println((responseObject.pokemon))

	for _, pokemon := range responseObject.pokemon {
		fmt.Println(pokemon.especie.nome)
	}
}