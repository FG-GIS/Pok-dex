package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/FG-GIS/bootpokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}

type Config struct {
	PokeApiClient pokeapi.Client // implement pokeapi client
	NextLoc       *string
	PrevLoc       *string
}

func cleanInput(text string) []string {
	cleanOut := strings.Fields(strings.ToLower(text))
	return cleanOut
}

func StartPokeCli(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		cmd, ok := GetCommands()[text[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
