package main

func (p *Parser) ParseLiteral(parent *Expression, token Token) error {
	Expr(parent, VARIANTS.Literal, token)

	return nil
}
