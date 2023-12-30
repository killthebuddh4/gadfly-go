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

type Environment struct {
	name   string
	parent *Environment
	values map[string]Value
}

var root = Environment{
	name:   "root",
	parent: nil,
	values: make(map[string]Value),
}

var env = &root

func PushEnvironment() {
	e := Environment{
		name:   "child",
		parent: env,
		values: make(map[string]Value),
	}

	env = &e
}

func PopEnvironment() {
	env = env.parent
}

func GetSymbol(name string) (Value, error) {
	return getSymbol(env, name)
}

func getSymbol(e *Environment, name string) (Value, error) {
	value, ok := e.values[name]

	if !ok {
		if e.parent == nil {
			return nil, errors.New("symbol not found")
		} else {
			return getSymbol(e.parent, name)
		}
	}

	return value, nil
}

func SetSymbol(name string, value Value) error {
	_, ok := env.values[name]

	if ok {
		return errors.New("symbol already exists")
	}

	env.values[name] = value

	return nil
}
