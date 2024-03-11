package types

import (
	"fmt"
	"strings"
)

type Schema func(value Value) (Value, error)

type Predicate func(lexeme Lexeme) bool

type Block = []*Expression

type Signature struct {
	Parameters map[string]string
	Returns    []*Expression
}

type Expression struct {
	Parent       *Expression
	Operator     Operator
	Signature    *Signature
	Arguments    map[string]Block
	Catches      []Block
	Trajectories []*Trajectory
}

func Print(expr *Expression, indent int) {
	if indent == 0 {
		fmt.Println("PRINTING PARSE TREE")
	}

	fmt.Println(strings.Repeat(" ", indent), expr.Operator.Value)

	for name, block := range expr.Arguments {
		fmt.Println(strings.Repeat(" ", indent+2), name)

		for _, exp := range block {
			Print(exp, indent+4)
		}
	}
}
