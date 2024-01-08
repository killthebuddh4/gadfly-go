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
		eval(args[1])
	}
}

func eval(pathToFile string) {
	source := ""

	files := []string{"lib.math.fly", "lib.array.fly", pathToFile}

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
		source += "\n"
		source += string(data)
	}

	fmt.Println("Source: ", source)

	SetSource(source)

	tokens, scanErr := Scan(GetSource())

	SetTokens(tokens)

	if scanErr != nil {
		fmt.Println("Error scanning: ", scanErr)
		return
	}

	root := Expr(nil, Token{
		Type:   "ROOT",
		Start:  0,
		Length: 0,
	})

	InitializeStdLib(&root)

	parseErr := Parse(&root, GetTokens())

	if parseErr != nil {
		fmt.Println("Error parsing: ", parseErr)
		return
	}

	_, evalErr := Evaluate(root)

	if evalErr != nil {
		fmt.Println("Error evaluating: ", evalErr)
		return
	}
}
