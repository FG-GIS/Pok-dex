package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	cleanOut := strings.Fields(strings.ToLower(text))
	return cleanOut
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}
		text := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %v\n", text[0])
		fmt.Print("Pokedex > ")
	}
}
