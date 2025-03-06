package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/FG-GIS/bootpokedex/internal/pokeapi"
)

func commandHelp(c *Config, arg string) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, v := range GetCommands() {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}

func commandExit(c *Config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	defer os.Exit(0)

	return nil
}

func commandMapF(c *Config, arg string) error {
	areas, err := c.PokeApiClient.GetLocations(c.NextLoc)
	if err != nil {
		return err
	}
	c.NextLoc = areas.Next
	c.PrevLoc = areas.Previous

	for _, loc := range areas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
func commandMapB(c *Config, arg string) error {
	if c.PrevLoc == nil {
		return errors.New("you're on the first page")
	}
	areas, err := c.PokeApiClient.GetLocations(c.PrevLoc)
	if err != nil {
		return err
	}
	c.NextLoc = areas.Next
	c.PrevLoc = areas.Previous

	for _, loc := range areas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(c *Config, arg string) error {
	locationDetail, err := c.PokeApiClient.ExploreLocation(arg)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", arg)
	fmt.Println("Found Pokemon:")
	for _, poke := range pokeapi.GetPokemons(locationDetail) {
		fmt.Printf(" - %v\n", poke)
	}
	return nil
}

func commandCatch(c *Config, arg string) error {
	pokeDetail, err := c.PokeApiClient.GetPokeDetail(arg)
	if err != nil {
		return err
	}
	pokeapi.Catch(pokeDetail)
	return nil
}

func commandInspect(c *Config, arg string) error {
	entry := pokeapi.Inspect(arg)
	if entry.Status != "" {
		fmt.Println(entry.Status)
		return nil
	}
	fmt.Printf("Name: %v\n", entry.Name)
	fmt.Printf("Height: %v\n", entry.Height)
	fmt.Printf("Weight: %v\n", entry.Weight)
	fmt.Printf("Stats:\n")
	fmt.Printf("  -hp: %v\n", entry.Stats["hp"])
	fmt.Printf("  -attack: %v\n", entry.Stats["attack"])
	fmt.Printf("  -defense: %v\n", entry.Stats["defense"])
	fmt.Printf("  -special-attack: %v\n", entry.Stats["special-attack"])
	fmt.Printf("  -special-defense: %v\n", entry.Stats["special-defense"])
	fmt.Printf("  -speed: %v\n", entry.Stats["speed"])
	fmt.Printf("Types:\n")
	for _, v := range entry.Types {
		fmt.Printf("  -%v\n", v)
	}
	return nil
}

func commandPokedex(c *Config, arg string) error {
	pkmList := pokeapi.ListPokedex()
	fmt.Printf("Your Pokedex:\n")
	for _, v := range pkmList {
		fmt.Printf(" - %v\n", v)
	}
	return nil
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Prints the list of commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Prints location-areas",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints location-areas previous page",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "explore <location>\n will list the pokemons on the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch <pokemon>\n will attempt to catch the speciman",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect <pokemon>\n retrieve info from pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "print out your pokedex (list of caught pokemons)",
			callback:    commandPokedex,
		},
	}
}
