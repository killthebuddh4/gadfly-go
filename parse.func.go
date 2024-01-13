package main

import (
	"errors"
)

func (p *Parser) ParseFunc(parent *Expression, operator Token) error {
	root := Expr(parent, VARIANTS.Lambda, operator)

	parameters := []string{}

	if accept(p, isPipe) {
		for accept(p, isIdentifier) {
			parameter := p.previous().Lexeme

			definition := Definition{
				Name:     parameter,
				Arity:    0,
				Variadic: false,
			}

			err := Define(&root, definition.Name, definition)

			if err != nil {
				return nil
			}

			parameters = append(parameters, parameter)
		}

		if !accept(p, isPipe) {
			return errors.New("expected closing pipe")
		}

		root.Parameters = parameters
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
			return nil
		}
	}

	for {
		if accept(p, isEnd) {
			break
		}

		if p.isAtEnd() {
			break
		}

		err := p.expression(&root)

		if err != nil {
			return nil
		}
	}

	return nil
}
