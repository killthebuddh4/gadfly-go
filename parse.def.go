package main

import (
	"errors"
)

func (p *Parser) ParseDef(parent *Expression, operator Token) error {
	root := Expr(parent, VARIANTS.Def, operator)

	if !accept(p, isIdentifier) {
		return errors.New("expected identifier after def")
	} else {
		p.ParseLiteral(&root)
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
			return err
		}
	}

	return nil
}
