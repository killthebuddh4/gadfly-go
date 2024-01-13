package main

import (
	"errors"
	"fmt"
)

func Resolve(expression *Expression, name string) (Definition, error) {
	if expression == nil {
		return Definition{}, errors.New("value not found for " + name)
	}

	for _, defn := range expression.Definitions {
		if defn.Name == name {
			return defn, nil
		}
	}

	return Resolve(expression.Parent, name)
}

func Define(expression *Expression, name string, defn Definition) error {
	if expression == nil {
		return errors.New("cannot define name in nil expression")
	}

	_, err := Resolve(expression, name)

	if err == nil {
		return errors.New("name " + name + " is already defined")
	}

	fmt.Println("Defining name <", name, "> in expression <", expression.Operator.Lexeme, ">")
	expression.Definitions = append(expression.Definitions, defn)

	return nil
}

func Override(expression *Expression, name string, defn Definition) error {
	if expression == nil {
		return errors.New("definition not found for " + name)
	}

	_, err := Resolve(expression, name)

	if err != nil {
		return errors.New("cannot override a definition that is not yet defined, tried: " + name)
	}

	expression.Definitions = append(expression.Definitions, defn)

	return nil
}
