package main

import (
	"fmt"
	"io"
	"os"
)

type Interpreter struct {
	source string
	lexed  []Token
	parsed Expression
}

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: gadflai [script]")
		fmt.Println(len(args))
	} else {
		pathToFile := args[1]

		fmt.Println("Running file: ", pathToFile)

		file, openFileErr := os.Open(pathToFile)

		if openFileErr != nil {
			fmt.Println("Error opening file: ", openFileErr)
			return
		}

		defer file.Close()

		data, readFileErr := io.ReadAll(file)

		if readFileErr != nil {
			fmt.Println("Error reading file: ", readFileErr)
			return
		}

		scanner := Scanner{
			Source:  string(data),
			Tokens:  []Token{},
			Start:   0,
			Current: 0,
		}

		tokens, scanErr := Scan(string(data))

		if scanErr != nil {
			fmt.Println("Error scanning: ", scanErr)
			return
		}

		for _, token := range tokens {
			PrintToken(scanner.Source, token)
		}

		expression, parseErr := Parse(tokens)

		if parseErr != nil {
			fmt.Println("Error parsing: ", parseErr)
			return
		}

		PrintExpression(string(data), expression)
	}
}
