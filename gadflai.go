package main

import (
	"fmt"
	"io"
	"os"
)

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

		SetSource(string(data))

		if readFileErr != nil {
			fmt.Println("Error reading file: ", readFileErr)
			return
		}

		tokens, scanErr := Scan(GetSource())

		SetTokens(tokens)

		if scanErr != nil {
			fmt.Println("Error scanning: ", scanErr)
			return
		}

		for _, token := range GetTokens() {
			PrintToken(token)
		}

		expression, parseErr := Parse(GetTokens())

		SetExpression(expression)

		if parseErr != nil {
			fmt.Println("Error parsing: ", parseErr)
			return
		}

		PrintExpression(GetSource(), GetExpression())

		val, evalErr := Evaluate(GetExpression())

		if evalErr != nil {
			fmt.Println("Error evaluating: ", evalErr)
			return
		}

		fmt.Println("Result: ", val)
	}
}
