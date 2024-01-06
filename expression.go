package main

type Expression struct {
	Parent   *Expression
	Operator Token
	Inputs   []Expression
	Values   map[string]Value
	Keywords map[string]string
}

func Expr(parent *Expression, operator Token) Expression {
	return Expression{
		Parent:   parent,
		Operator: operator,
		Inputs:   []Expression{},
		Values:   make(map[string]Value),
		Keywords: make(map[string]string),
	}
}
