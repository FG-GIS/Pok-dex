package cli

import (
	"errors"
	"fmt"
	"os"
)

func commandHelp(c *Config) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, v := range GetCommands() {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	defer os.Exit(0)

	return nil
}

func commandMapF(c *Config) error {
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
func commandMapB(c *Config) error {
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
	}
}
