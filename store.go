package main

var source *string

func GetSource() string {
	return *source
}

func SetSource(s string) {
	source = &s
}

var tokens *[]Token

func GetTokens() []Token {
	return *tokens
}

func SetTokens(t []Token) {
	tokens = &t
}

var program *[]Expression

func GetProgram() []Expression {
	return *program
}

func SetProgram(p []Expression) {
	program = &p
}

var symbols = make(map[string]Value)

func GetSymbol(name string) Value {
	return symbols[name]
}

func SetSymbol(name string, value Value) {
	symbols[name] = value
}
