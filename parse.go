package main

import (
	"errors"
)

func Parse(root *Expression, tokens []Token) error {
	if root == nil {
		return errors.New("cannot parse nil expression")
	}

	parser := Parser{
		Tokens:  tokens,
		Current: 0,
	}

	for !parser.isAtEnd() {
		err := parser.expression(root)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) expression(parent *Expression) error {
	if accept(p, isDef) {
		return p.ParseDef(parent, p.previous())
	} else if accept(p, isFn) {
		return p.ParseFunc(parent, p.previous())
	} else if accept(p, isIdentifier) {
		return p.ParseIdentifier(parent, p.previous())
	} else {
		return p.ParseCalc(parent)
	}
}
