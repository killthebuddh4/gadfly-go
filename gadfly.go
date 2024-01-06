package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	InitializeStdLib()

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: gadflai [script]")
		fmt.Println(len(args))
	} else {
		eval(args[1])
	}
}

func eval(pathToFile string) {
	source := ""

	// files := []string{"lib.math.gf", "lib.array.gf", pathToFile}
	files := []string{pathToFile}

	for _, file := range files {
		f, err := os.Open(file)

		if err != nil {
			fmt.Println("Error opening file: ", err)
			return
		}

		defer f.Close()

		data, err := io.ReadAll(f)

		if err != nil {
			fmt.Println("Error reading file: ", err)
			return
		}

		source += "\n"
		source += string(data)
	}

	SetSource(source)

	tokens, scanErr := Scan(GetSource())

	SetTokens(tokens)

	if scanErr != nil {
		fmt.Println("Error scanning: ", scanErr)
		return
	}

	expression, parseErr := Parse(GetTokens())

	SetProgram(expression)

	if parseErr != nil {
		fmt.Println("Error parsing: ", parseErr)
		return
	}

	values := []Value{}

	for _, expression := range GetProgram() {
		val, evalErr := Evaluate(expression)

		if evalErr != nil {
			fmt.Println("Error evaluating: ", evalErr)
			return
		}

		values = append(values, val)
	}
}
