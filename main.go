package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Split(strings.ToLower(text), " ")
}

func main() {
	fmt.Println(cleanInput("Hello, World!"))
}
