package main

import "errors"

func Expr(parent *Expression, variant string, operator Token) Expression {
	expr := Expression{
		Parent:       parent,
		Operator:     operator,
		Variant:      variant,
		Children:     []*Expression{},
		Parameters:   []string{},
		Trajectories: []*Trajectory{},
		Definitions:  []Definition{},
	}

	if parent != nil {
		parent.Children = append(parent.Children, &expr)
	}

	return expr
}

func RootExpr() Expression {
	return Expression{
		Parent: nil,
		Operator: Token{
			Type:   "ROOT",
			Lexeme: "ROOT",
			Start:  0,
			Length: 0,
		},
		Children:     []*Expression{},
		Parameters:   []string{},
		Trajectories: []*Trajectory{},
		Definitions:  BUILTINS,
	}
}

func accept(p *Parser, predicate Predicate) bool {
	token := p.read()
	if predicate(token) {
		p.advance()
		return true
	} else {
		return false
	}
}

func (p *Parser) advance() error {
	if p.isAtEnd() {
		return errors.New("unexpected end of file")
	}

	p.Current++

	return nil
}

func (p Parser) read() Token {
	return p.Tokens[p.Current]
}

func (p Parser) previous() Token {
	return p.Tokens[p.Current-1]
}

func (p Parser) isAtEnd() bool {
	return p.Current >= len(p.Tokens)-1
}
