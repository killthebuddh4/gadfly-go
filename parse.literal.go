package main

import "errors"

func (p *Parser) ParseLiteral(parent *Expression, token Token) error {
	if token.Type != TOKENS.String && token.Type != TOKENS.Number {
		return errors.New("expected literal token")
	}

	Expr(parent, VARIANTS.Literal, token)
	return nil
}
