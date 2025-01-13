package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		fmt.Println("echoing: ", cleaned)
	}
}

func cleanInput(str string) []string { // parse the input string and clean it up and return all the tokens in the string
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
