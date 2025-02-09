package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			input := scanner.Text()

			if input == "exit" {
				os.Exit(1)
			}

			first := cleanInput(input)[0]
			fmt.Printf("Your command was: %s\n", first)
			fmt.Print("Pokedex > ")

		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		}
	}
}
