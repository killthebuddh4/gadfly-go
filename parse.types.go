package main

type Parser struct {
	Tokens  []Token
	Current int
}

type Definition struct {
	// Name of the function
	Name string
	// number of required parameters
	Arity int
	// whether to allow Variadic parameters
	Variadic bool
}

type Variants struct {
	Lambda  string
	Call    string
	Literal string
}

var VARIANTS = Variants{
	Lambda:  "lambda",
	Call:    "call",
	Literal: "literal",
}

type Expression struct {
	Parent       *Expression
	Variant      string
	Operator     Token
	Parameters   []string
	Children     []*Expression
	Trajectories []*Trajectory
	Definitions  []Definition
}

type Predicate func(token Token) bool
