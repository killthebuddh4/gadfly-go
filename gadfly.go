package main

import (
	"fmt"
	"io"
	"os"

	eval1 "github.com/killthebuddh4/gadflai/eval"
	"github.com/killthebuddh4/gadflai/lex"
	"github.com/killthebuddh4/gadflai/parse"
	"github.com/killthebuddh4/gadflai/types"
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

	files := []string{"lib.math.fly", "lib.array.fly", "lib.map.fly", pathToFile}

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

	lexemes, err := lex.Lex(source)

	if err != nil {
		fmt.Println("Error lexing: ", err)
		return
	}

	rootOperator, err := types.NewOperator("program")

	if err != nil {
		fmt.Println("Error creating root operator: ", err)
		return
	}

	rootExp := types.NewExpression(nil, rootOperator, []*types.Expression{})

	parseErr := parse.Parse(&rootExp, lexemes)

	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		types.Print(rootExp, 0)
	}

	root := types.NewTrajectory(nil, &rootExp)

	if parseErr != nil {
		fmt.Println("Error parsing: ", parseErr)
		return
	}

	_, err = eval1.Eval(&root)

	if err != nil {
		fmt.Println("Error evaluating: ", err)
		return
	}

}
