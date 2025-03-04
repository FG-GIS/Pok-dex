package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//STRUCTS

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}
type config struct {
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []pokeMap `json:"results"`
}

type pokeMap struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

//CONSTS

const (
	POKEAPI string = "https://pokeapi.co/api/v2/"
	AREAS   string = "location-area"
)

//VARS

var validCommands map[string]cliCommand
var mapConfig = config{}

//FUNCS

func cleanInput(text string) []string {
	cleanOut := strings.Fields(strings.ToLower(text))
	return cleanOut
}

func commandMap(c *config) error {
	url := POKEAPI + AREAS
	if c.Next == "" && c.Previous != "" {
		fmt.Println("You've reached the end of the list")
		return nil
	}
	if c.Next != "" {
		url = c.Next
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error in the Areas request: %w", err)
	}
	defer res.Body.Close()

	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error retrieving the payload: %w", err)
	}
	err = json.Unmarshal(payload, &mapConfig)
	if err != nil {
		return fmt.Errorf("error decoding the payload: %w", err)
	}

	for _, pm := range mapConfig.Results {
		fmt.Println(pm.Name)
	}

	return nil
}

func commandMapB(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area?offset=1069&limit=20"
	if c.Next != "" && c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	if c.Previous != "" {
		url = c.Previous
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error in the Areas request: %w", err)
	}
	defer res.Body.Close()

	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error retrieving the payload: %w", err)
	}
	err = json.Unmarshal(payload, &mapConfig)
	if err != nil {
		return fmt.Errorf("error decoding the payload: %w", err)
	}

	for _, pm := range mapConfig.Results {
		fmt.Println(pm.Name)
	}

	return nil
}

func commandHelp(c *config) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, v := range validCommands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	defer os.Exit(0)

	return fmt.Errorf("user interrupt - exit command")
}

//LOGIC

func main() {
	validCommands = map[string]cliCommand{
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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints location-areas previous page",
			callback:    commandMapB,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}
		text := cleanInput(scanner.Text())
		cmd, ok := validCommands[text[0]]
		if !ok {
			fmt.Println("Unknown command")
		}
		err := cmd.callback(&mapConfig)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("Pokedex > ")
	}
}
