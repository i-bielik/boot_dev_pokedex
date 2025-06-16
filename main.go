package main

import (
	"time"

	pokeapi "github.com/i-bielik/pokedexcli/internal"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)

}
