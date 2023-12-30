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
