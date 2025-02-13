package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand) func(*Config) error {
	return func(cfg *Config) error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")

		for _, c := range commands {
			fmt.Printf("%s: %s\n", c.name, c.description)
		}
		return nil
	}
}

func commandMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.NextUrl != "" {
		url = cfg.NextUrl
	}

	res, err := GetPokeLocations(url)
	if err != nil {
		return err
	}

	locations := Locations{}
	jsonErr := json.Unmarshal(res, &locations)
	if jsonErr != nil {
		return err
	}

	cfg.NextUrl = locations.Next
	cfg.PrevUrl = locations.Prev

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapBack(cfg *Config) error {
	var url string
	if cfg.PrevUrl != "" {
		url = cfg.PrevUrl
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := GetPokeLocations(url)
	if err != nil {
		return err
	}

	locations := Locations{}
	jsonErr := json.Unmarshal(res, &locations)
	if jsonErr != nil {
		return err
	}

	cfg.NextUrl = locations.Next
	cfg.PrevUrl = locations.Prev

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Locations struct {
	Count   int      `json:"count"`
	Next    string   `json:"next"`
	Prev    string   `json:"previous"`
	Results []Result `json:"results"`
}

type Result struct {
	Name string
	Url  string
}

type Config struct {
	NextUrl string
	PrevUrl string
}

func main() {
	cliMap := map[string]cliCommand{}
	cfg := &Config{}

	cliMap["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	cliMap["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp(cliMap),
	}

	cliMap["map"] = cliCommand{
		name:        "map",
		description: "Prints 20 locations, each consecutive run of 'map' will print the next set of locations",
		callback:    commandMap,
	}

	cliMap["mapb"] = cliCommand{
		name:        "mapb",
		description: "Goes back to the last printed list of 20 locations",
		callback:    commandMapBack,
	}

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			input := scanner.Text()
			cleaned := cleanInput(input)

			if len(cleaned) > 0 {
				command := cleaned[0]
				if cmd, ok := cliMap[command]; ok {
					err := cmd.callback(cfg)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println("Unknown command")
				}
			}
			fmt.Print("Pokedex > ")

		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		}
	}
}
