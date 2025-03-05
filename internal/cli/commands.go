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
	}
}
