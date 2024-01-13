package main

import (
	"errors"
	"fmt"
)

func (p *Parser) ParseFunc(parent *Expression, operator Token) error {
	fmt.Println("Going to parse a func, lexeme is <", p.previous().Lexeme, ">")

	root := Expr(parent, VARIANTS.Call, operator)

	parameters := []string{}

	if accept(p, isPipe) {
		fmt.Println("Going to parse parameters, lexeme is <", p.previous().Lexeme, ">")

		for accept(p, isIdentifier) {
			fmt.Println("Going to parse one parameter, lexeme is <", p.previous().Lexeme, ">")

			parameter := p.previous().Lexeme

			definition := Definition{
				Name:     parameter,
				Arity:    0,
				Variadic: false,
			}

			err := Define(root, definition.Name, definition)

			if err != nil {
				return nil
			}

			parameters = append(parameters, parameter)
		}

		if !accept(p, isPipe) {
			return errors.New("expected closing pipe")
		}

		root.Parameters = parameters

		fmt.Println("Done parsing parameters, found <", len(parameters), "> parameters")
	}

	if parent.Operator.Lexeme == "def" {
		lexeme := parent.Children[0].Operator.Lexeme

		definition := Definition{
			Name:     lexeme,
			Arity:    len(root.Parameters),
			Variadic: false,
		}

		err := Define(parent.Parent, definition.Name, definition)

		if err != nil {
			return err
		}
	}

	for {
		if accept(p, isEnd) {
			break
		}

		if p.isAtEnd() {
			break
		}

		err := p.expression(root)

		if err != nil {
			return nil
		}
	}

	fmt.Println("FN PARENT PARENT TREE")
	PrintExp(parent.Parent)
	fmt.Println("FN PARENT TREE")
	PrintExp(parent)

	fmt.Println("Done parsing func with N children: ", len(root.Children))

	return nil
}
