package main

import (
	"errors"
	"fmt"
)

type Lambda func(args ...Value) (Value, error)

type Expression struct {
	Parent      *Expression
	Operator    Token
	Parameters  []string
	Children    []*Expression
	Definitions map[string]Value
}

func Expr(parent *Expression, operator Token) Expression {
	return Expression{
		Parent:      parent,
		Operator:    operator,
		Children:    []*Expression{},
		Parameters:  []string{},
		Definitions: make(map[string]Value),
	}
}

func PrintExp(exp *Expression, indent int) {
	if exp == nil {
		return
	}

	ind := func() {
		for i := 0; i < indent; i++ {
			fmt.Print("    ")
		}
	}

	ind()
	fmt.Println(exp.Operator.Lexeme)

	for key := range exp.Definitions {
		ind()
		fmt.Printf("%s: %v\n", key, exp.Definitions[key])
	}

	for _, child := range exp.Children {
		PrintExp(child, indent+4)
	}
}

func ResolveDef(inExp *Expression, name string) (Value, error) {
	if inExp == nil {
		return nil, errors.New("definition not found for " + name)
	}

	for kw, val := range inExp.Definitions {
		if kw == name {
			return val, nil
		}
	}

	return ResolveDef(inExp.Parent, name)
}

func DefineDef(inExp *Expression, name string, val Value) error {
	if inExp == nil {
		return errors.New("cannot define value in nil expression")
	}

	_, ok := inExp.Definitions[name]

	if ok {
		return errors.New("definition already exists for name " + name)
	}

	inExp.Definitions[name] = val

	return nil
}

func EditDef(inExp *Expression, name string, val Value) error {
	if inExp == nil {
		return errors.New("definition not found for " + name)
	}

	for kw := range inExp.Definitions {
		if kw == name {
			inExp.Definitions[name] = val
			return nil
		}
	}

	return EditDef(inExp.Parent, name, val)
}

func ClearDefs(inExp *Expression) error {
	if inExp == nil {
		return errors.New("cannot clear definitions in nil expression")
	}

	inExp.Definitions = make(map[string]Value)

	return nil
}
