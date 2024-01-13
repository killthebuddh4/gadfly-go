package main

import (
	"errors"
)

type Parser struct {
	Tokens  []Token
	Current int
}

func Parse(root *Expression, tokens []Token) error {
	parser := Parser{
		Tokens:  tokens,
		Current: 0,
	}

	return parser.program(root)
}

func (p *Parser) program(root *Expression) error {
	for !p.isAtEnd() {
		err := p.expression(root)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) expression(parent *Expression) error {
	if parent == nil {
		return errors.New("cannot parse expression with nil parent")
	}

	if accept(p, isFn) {
		child, err := p.fn(parent)

		if err != nil {
			return err
		}

		child.Parent = parent
		parent.Children = append(parent.Children, child)
	} else if accept(p, isBlockStart) {
		child, err := p.block(parent)

		if err != nil {
			return err
		}

		child.Parent = parent
		parent.Children = append(parent.Children, child)
	} else {
		child, err := p.logical(parent)

		if err != nil {
			return err
		}

		child.Parent = parent
		parent.Children = append(parent.Children, child)
	}

	return nil
}

func (p *Parser) fn(parent *Expression) (*Expression, error) {
	operator := p.previous()

	root := Expr(parent, operator)

	if accept(p, isPipe) {
		parameters := []string{}

		for accept(p, isIdentifier) {
			parameters = append(parameters, p.previous().Lexeme)
		}

		if !accept(p, isPipe) {
			return nil, errors.New("expected closing pipe")
		}

		root.Parameters = parameters
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
			return nil, err
		}
	}

	return &root, nil
}

func (p *Parser) block(parent *Expression) (*Expression, error) {
	operator := p.previous()

	root := Expr(parent, operator)

	for {
		if accept(p, isEnd) {
			break
		}

		if p.isAtEnd() {
			break
		}

		err := p.expression(&root)

		if err != nil {
			return nil, err
		}
	}

	return &root, nil
}

func (p *Parser) logical(parent *Expression) (*Expression, error) {
	left, err := p.equality(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isLogical) {
		operator := p.previous()

		right, err := p.equality(parent)

		if err != nil {
			return nil, err
		}

		left = &Expression{
			Operator: operator,
			Children: []*Expression{left, right},
		}
	}

	return left, nil
}

func (p *Parser) equality(parent *Expression) (*Expression, error) {
	left, err := p.comparison(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isEquality) {
		operator := p.previous()

		right, err := p.comparison(parent)

		if err != nil {
			return nil, err
		}

		left = &Expression{
			Operator: operator,
			Children: []*Expression{left, right},
		}
	}

	return left, nil
}

func (p *Parser) comparison(parent *Expression) (*Expression, error) {
	left, err := p.term(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isComparison) {
		operator := p.previous()

		right, err := p.term(parent)

		if err != nil {
			return nil, err
		}

		left = &Expression{
			Operator: operator,
			Children: []*Expression{left, right},
		}
	}

	return left, nil
}

func (p *Parser) term(parent *Expression) (*Expression, error) {
	left, err := p.factor(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isTerm) {
		operator := p.previous()

		right, err := p.factor(parent)

		if err != nil {
			return nil, err
		}

		left = &Expression{
			Operator: operator,
			Children: []*Expression{left, right},
		}
	}

	return left, nil
}

func (p *Parser) factor(parent *Expression) (*Expression, error) {
	left, err := p.unary(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isFactor) {
		operator := p.previous()

		right, err := p.unary(parent)

		if err != nil {
			return nil, err
		}

		left = &Expression{
			Operator: operator,
			Children: []*Expression{left, right},
		}
	}

	return left, nil
}

func (p *Parser) unary(parent *Expression) (*Expression, error) {
	if accept(p, isUnary) {
		operator := p.previous()

		right, err := p.unary(parent)

		if err != nil {
			return nil, err
		}

		return &Expression{
			Operator: operator,
			Children: []*Expression{right},
		}, nil
	}

	return p.atom(parent)
}

func (p *Parser) atom(parent *Expression) (*Expression, error) {
	if accept(p, isAtom) {
		operator := p.previous()

		result := Expr(parent, operator)

		return &result, nil
	}

	return nil, errors.New("expected expression but got " + p.read().Type)
}

type Predicate func(token Token) bool

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
