package main

import (
	"errors"
	"fmt"
)

func (p *Parser) ParseLogic(parent *Expression) error {
	fmt.Println("Going to parse a logic, lexeme is <", p.previous().Lexeme, ">")

	left, err := p.equality(parent)

	if err != nil {
		return err
	}

	for accept(p, isLogical) {
		fmt.Println("Going to parse and or or, lexeme is <", p.previous().Lexeme, ">")

		operator := p.previous()

		right, err := p.equality(parent)

		if err != nil {
			return err
		}

		left = &Expression{
			Operator: operator,
			Variant:  VARIANTS.Operator,
			Children: []*Expression{left, right},
		}
	}

	parent.Children = append(parent.Children, left)

	fmt.Println("Done parsing logic with N children: ", len(parent.Children))

	return nil
}

func (p *Parser) equality(parent *Expression) (*Expression, error) {
	left, err := p.comparison(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isEquality) {
		fmt.Println("Going to parse equality, lexeme is <", p.previous().Lexeme, ">")
		operator := p.previous()

		right, err := p.comparison(parent)

		if err != nil {
			return nil, err
		}

		left = &Expression{
			Operator: operator,
			Variant:  VARIANTS.Operator,
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
		fmt.Println("Going to parse comparison, lexeme is <", p.previous().Lexeme, ">")
		operator := p.previous()

		right, err := p.term(parent)

		if err != nil {
			return nil, err
		}

		left = &Expression{
			Operator: operator,
			Variant:  VARIANTS.Operator,
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
		fmt.Println("Going to parse term, lexeme is <", p.previous().Lexeme, ">")
		operator := p.previous()

		right, err := p.factor(parent)

		if err != nil {
			return nil, err
		}

		left = &Expression{
			Operator: operator,
			Variant:  VARIANTS.Operator,
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
		fmt.Println("Going to parse factor, lexeme is <", p.previous().Lexeme, ">")
		operator := p.previous()

		right, err := p.unary(parent)

		if err != nil {
			return nil, err
		}

		left = &Expression{
			Operator: operator,
			Variant:  VARIANTS.Operator,
			Children: []*Expression{left, right},
		}
	}

	return left, nil
}

func (p *Parser) unary(parent *Expression) (*Expression, error) {
	if accept(p, isUnary) {
		fmt.Println("Going to parse unary, lexeme is <", p.previous().Lexeme, ">")

		operator := p.previous()

		right, err := p.unary(parent)

		if err != nil {
			return nil, err
		}

		return &Expression{
			Operator: operator,
			Variant:  VARIANTS.Operator,
			Children: []*Expression{right},
		}, nil
	}

	return p.atom(parent)
}

func (p *Parser) atom(parent *Expression) (*Expression, error) {
	if accept(p, isAtom) {
		fmt.Println("Going to parse atom, lexeme is <", p.previous().Lexeme, ">")

		operator := p.previous()

		switch operator.Type {
		case TOKENS.Identifier:
			return &Expression{
				Operator: operator,
				Variant:  VARIANTS.Call,
				Children: []*Expression{},
			}, nil
		case TOKENS.String, TOKENS.False, TOKENS.True, TOKENS.Number, TOKENS.Nil:
			return &Expression{
				Operator: operator,
				Variant:  VARIANTS.Literal,
				Children: []*Expression{},
			}, nil
		default:
			return nil, errors.New("expected atom but got " + operator.Type)
		}
	}

	return nil, errors.New("expected expression but got " + p.read().Type)
}
