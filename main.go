package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var validCommands map[string]cliCommand

func cleanInput(text string) []string {
	cleanOut := strings.Fields(strings.ToLower(text))
	return cleanOut
}

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, v := range validCommands {
		fmt.Printf("%v: %v\n", v.name, v.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	defer os.Exit(0)

	return fmt.Errorf("user interrupt - exit command")
}

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
		err := cmd.callback()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("Pokedex > ")
	}
}
