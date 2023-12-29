package main

import "fmt"

type Expression struct {
	Operator Token
	Inputs   []Expression
}

func PrintExpression(source string, exp Expression) {
	fmt.Println("----------------------------------------------")
	fmt.Println("OPERATOR")

	PrintToken(source, exp.Operator)

	fmt.Println("INPUTS")

	for _, input := range exp.Inputs {
		PrintToken(source, input.Operator)
	}

	fmt.Println("----------------------------------------------")

	for _, input := range exp.Inputs {
		PrintExpression(source, input)
	}
}
