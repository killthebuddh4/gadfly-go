package main

import (
	"fmt"
	"io"
	"os"

	"github.com/killthebuddh4/gadflai/exec"
	exp "github.com/killthebuddh4/gadflai/expression"
	"github.com/killthebuddh4/gadflai/lex"
	"github.com/killthebuddh4/gadflai/parse"
	traj "github.com/killthebuddh4/gadflai/trajectory"
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

	files := []string{"lib.math.fly", "lib.array.fly", "lib.record.fly", pathToFile}

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

	rootOperator, err := exp.NewOperator("program")

	if err != nil {
		fmt.Println("Error creating root operator: ", err)
		return
	}

	rootExp := exp.NewExpression(nil, rootOperator, []*exp.Expression{})

	parseErr := parse.Parse(&rootExp, lexemes)

	_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

	if debug {
		exp.Print(rootExp, 0)
	}

	root := traj.Traj(nil, &rootExp)

	if parseErr != nil {
		fmt.Println("Error parsing: ", parseErr)
		return
	}

	_, err = exec.Exec(&root)

	if err != nil {
		fmt.Println("Error evaluating: ", err)
		return
	}

}
