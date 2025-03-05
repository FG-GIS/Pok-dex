package main

import (
	"time"

	"github.com/FG-GIS/bootpokedex/internal/cli"
	"github.com/FG-GIS/bootpokedex/internal/pokeapi"
)

func main() {
	pokeC := pokeapi.NewClient(time.Second*5, time.Minute)
	config := &cli.Config{
		PokeApiClient: pokeC,
	}
	cli.StartPokeCli(config)
}
