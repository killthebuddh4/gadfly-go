package main

func (p *Parser) ParseLiteral(parent *Expression) error {
	Expr(parent, VARIANTS.Literal, p.previous())
	return nil
}
