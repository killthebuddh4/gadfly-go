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

	rootOperator := types.Operator{
		Type:  "program",
		Value: "program",
	}

	if err != nil {
		fmt.Println("Error creating root operator: ", err)
		return
	}

	rootExp := types.Expression{
		Parent:   nil,
		Operator: rootOperator,
		Signature: &types.Signature{
			Parameters: make(map[string]string),
			Returns:    []*types.Expression{},
		},
		Arguments:    map[string][]*types.Expression{},
		Catches:      []types.Block{},
		Trajectories: []*types.Trajectory{},
	}

	parseErr := parse.Parse(&rootExp, lexemes)

	root := types.NewTrajectory(nil, &rootExp)

	sString, err := eval.SchemaString(&root)

	if err != nil {
		fmt.Println("Error creating string schema: ", err)
		return
	}

	sNumber, err := eval.SchemaNumber(&root)

	if err != nil {
		fmt.Println("Error creating number schema: ", err)
		return
	}

	sBoolean, err := eval.SchemaBoolean(&root)

	if err != nil {
		fmt.Println("Error creating boolean schema: ", err)
		return
	}

	sArray, err := eval.SchemaArray(&root)

	if err != nil {
		fmt.Println("Error creating array schema: ", err)
		return
	}

	sHash, err := eval.SchemaHash(&root)

	if err != nil {
		fmt.Println("Error creating hash schema: ", err)
		return
	}

	sFunction, err := eval.SchemaFunction(&root)

	if err != nil {
		fmt.Println("Error creating function schema: ", err)
		return
	}

	sIdentity, err := eval.SchemaIdentity(&root)

	if err != nil {
		fmt.Println("Error creating identity schema: ", err)
		return
	}

	types.DefineName(&root, "String", sString)
	types.DefineName(&root, "Number", sNumber)
	types.DefineName(&root, "Boolean", sBoolean)
	types.DefineName(&root, "Array", sArray)
	types.DefineName(&root, "Hash", sHash)
	types.DefineName(&root, "Function", sFunction)
	types.DefineName(&root, "Identity", sIdentity)

	if parseErr != nil {
		fmt.Println("Error parsing: ", parseErr)
		return
	}

	_, err = eval.Exec(nil, &root, root.Expression)

	if err != nil {
		fmt.Println("Gadfly :: Error evaluating: ", err)
		return
	}

}
