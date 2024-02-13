package main

import (
	"fmt"
	"io"
	"os"

	"github.com/killthebuddh4/gadflai/eval"
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
		exec(args[1])
	}
}

func exec(pathToFile string) {
	source := ""

	_, excludeLib := os.LookupEnv("GADFLY_EXCLUDE_LIB")

	lib := []string{"lib.math.fly", "lib.array.fly", "lib.map.fly", "lib.schema.fly"}

	files := []string{pathToFile}

	if !excludeLib {
		files = append(lib, files...)
	}

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

	rootOperator, err := types.NewOperator("program", false)

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

	types.DefineName(&root, "String", eval.SchemaString())
	types.DefineName(&root, "Number", eval.SchemaNumber())
	types.DefineName(&root, "Boolean", eval.SchemaBoolean())
	types.DefineName(&root, "Array", eval.SchemaArray())
	types.DefineName(&root, "Hash", eval.SchemaHash())
	types.DefineName(&root, "Function", eval.SchemaFunction())
	types.DefineName(&root, "Identity", eval.SchemaIdentity())

	if parseErr != nil {
		fmt.Println("Error parsing: ", parseErr)
		return
	}

	_, err = eval.Exec(&root)

	if err != nil {
		fmt.Println("Error evaluating: ", err)
		return
	}

}
