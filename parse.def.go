package main

import (
	"errors"
	"fmt"
)

func (p *Parser) ParseDef(parent *Expression, operator Token) error {
	fmt.Println("Going to parse a def, lexeme is <", p.previous().Lexeme, ">")

	root := Expr(parent, VARIANTS.Call, operator)

	if !accept(p, isIdentifier) {
		return errors.New("expected identifier after def")
	} else {
		// Is this bad or is it fine?
		operator := p.previous()
		operator.Type = TOKENS.String
		err := p.ParseLiteral(root, operator)

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
			return err
		}
	}

	fmt.Println("Done parsing def with N children: ", len(root.Children))

	return nil
}
