package main

import "fmt"

func (p *Parser) ParseIdentifier(parent *Expression, operator Token) error {
	fmt.Println("Going to parse an identifier, lexeme is <", p.previous().Lexeme, ">")

	root := Expr(parent, VARIANTS.Call, operator)

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

	fmt.Println("Done parsing identifier <"+operator.Lexeme+"> with N children: ", len(root.Children))

	return nil
}
