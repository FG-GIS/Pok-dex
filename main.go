package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/FG-GIS/bootpokedex/internal/pokecache"
)

//STRUCTS

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}
type config struct {
	Actual   string
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
var mapConfig = config{
	Actual: POKEAPI + AREAS,
}
var mapCache = pokecache.NewCache(time.Second * 5)

//FUNCS

func cleanInput(text string) []string {
	cleanOut := strings.Fields(strings.ToLower(text))
	return cleanOut
}

func commandMap(c *config) error {
	var payload []byte

	// debug
	fmt.Println("Next     ==> ", c.Next)
	fmt.Println("Previous ==> ", c.Previous)

	if (c.Next == "" || c.Next == c.Actual) && c.Previous != "" {
		fmt.Println("You've reached the end of the list")
		return nil
	}
	if c.Next != "" {
		c.Actual = c.Next
	}
	fmt.Println("c.Actual      ==> ", c.Actual)

	if cData, ok := mapCache.Get(c.Actual); ok {
		// fmt.Println("test 1")
		payload = cData
	} else {
		res, err := http.Get(c.Actual)
		if err != nil {
			return fmt.Errorf("error in the Areas request: %w", err)
		}
		defer res.Body.Close()

		payload, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error retrieving the payload: %w", err)
		}
		// fmt.Println("test 1")
		mapCache.Add(c.Actual, payload)
	}

	err := json.Unmarshal(payload, &mapConfig)
	if err != nil {
		return fmt.Errorf("error decoding the payload: %w", err)
	}

	for _, pm := range mapConfig.Results {
		fmt.Println(pm.Name)
	}

	return nil
}

func commandMapB(c *config) error {
	var payload []byte

	// debug
	fmt.Println("Next     ==> ", c.Next)
	fmt.Println("Previous ==> ", c.Previous)

	if c.Next != "" && (c.Previous == "" || c.Previous == c.Actual) {
		fmt.Println("you're on the first page")
		return nil
	}
	if c.Previous != "" {
		c.Actual = c.Previous
	}
	fmt.Println("c.Actual      ==> ", c.Actual)

	if cData, ok := mapCache.Get(c.Actual); ok {
		// debug
		// fmt.Println("payload <== cache")

		payload = cData
	} else {
		res, err := http.Get(c.Actual)
		if err != nil {
			return fmt.Errorf("error in the Areas request: %w", err)
		}
		defer res.Body.Close()

		// debug
		// fmt.Println("payload <== new Data")

		payload, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error retrieving the payload: %w", err)
		}
		// debug
		// fmt.Println("cache <== new Data")
		// fmt.Println("Data ==>:\n", string(payload))
		mapCache.Add(c.Actual, payload)
	}
	// debug
	// fmt.Println("Data unmarshaling")
	// fmt.Println("Data ==>:\n", string(payload))
	err := json.Unmarshal(payload, &mapConfig)
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
			fmt.Print("Pokedex > ")
			continue
		}
		err := cmd.callback(&mapConfig)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("Pokedex > ")
	}
}
