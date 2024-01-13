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
	//
	Body Lambda
}

type Variants struct {
	Root     string
	Call     string
	Literal  string
	Operator string
}

var VARIANTS = Variants{
	Root:     "root",
	Call:     "call",
	Literal:  "literal",
	Operator: "operator",
}

type Expression struct {
	Id           string
	Parent       *Expression
	Variant      string
	Operator     Token
	Parameters   []string
	Children     []*Expression
	Trajectories []*Trajectory
	Definitions  []Definition
}

type Predicate func(token Token) bool
