package main

type Expression struct {
	Operator Token
	Inputs   []Expression
}
