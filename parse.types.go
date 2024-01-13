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
	Def     string
	Func    string
	Call    string
	Calc    string
	Literal string
}

var VARIANTS = Variants{
	Def:     "def",
	Func:    "func",
	Call:    "call",
	Calc:    "calc",
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
