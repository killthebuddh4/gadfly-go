package main

import (
	"errors"
)

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

func GetSymbol(name string) (Value, error) {
	value, ok := symbols[name]

	if !ok {
		return nil, errors.New("symbol not found")
	}

	return value, nil
}

func SetSymbol(name string, value Value) error {
	_, ok := symbols[name]

	if ok {
		return errors.New("symbol already exists")
	}

	symbols[name] = value

	return nil
}
