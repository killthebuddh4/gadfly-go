package main

type Lambda func(args ...Value) (Value, error)

type Expression struct {
	Parent      *Expression
	Operator    Token
	Inputs      []Expression
	Values      map[string]Value
	Definitions map[string]Lambda
}

func Expr(parent *Expression, operator Token) Expression {
	return Expression{
		Parent:      parent,
		Operator:    operator,
		Inputs:      []Expression{},
		Values:      make(map[string]Value),
		Definitions: make(map[string]Lambda),
	}
}
