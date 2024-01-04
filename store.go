package main

import (
	"errors"
	"fmt"
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
	val, err := getSymbol(env, name)

	if err != nil {
		PrintEnvironment()
		return nil, err
	}

	return val, nil
}

func getSymbol(e *Environment, name string) (Value, error) {
	value, ok := e.values[name]

	if !ok {
		if e.parent == nil {
			return nil, errors.New("symbol not found " + name)
		} else {
			return getSymbol(e.parent, name)
		}
	}

	return value, nil
}

func DefSymbol(name string, value Value) error {
	_, ok := env.values[name]

	if ok {
		return errors.New("symbol already exists")
	}

	env.values[name] = value

	return nil
}

func SetSymbol(name string, value Value) error {
	env.values[name] = value

	return nil
}

func PrintEnvironment() {
	e := env

	for e != nil {
		println("Environment: " + e.name)

		for k, v := range e.values {
			println("  " + k + " = " + fmt.Sprintf("%v", v))
		}

		e = e.parent
	}
}
