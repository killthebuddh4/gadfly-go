package main

import (
	"errors"
)

func (p *Parser) ParseCalc(parent *Expression) error {
	child, err := p.logical(parent)

	if err != nil {
		return err
	}

	parent.Children = append(parent.Children, child)
	child.Parent = parent

	return nil
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
			Variant:  VARIANTS.Call,
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
			Variant:  VARIANTS.Call,
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
			Variant:  VARIANTS.Call,
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
			Variant:  VARIANTS.Call,
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
			Variant:  VARIANTS.Call,
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
			Variant:  VARIANTS.Call,
			Children: []*Expression{right},
		}, nil
	}

	return p.atom(parent)
}

func (p *Parser) atom(parent *Expression) (*Expression, error) {
	if accept(p, isAtom) {
		operator := p.previous()

		result := Expr(parent, VARIANTS.Call, operator)

		return &result, nil
	}

	return nil, errors.New("expected expression but got " + p.read().Type)
}
