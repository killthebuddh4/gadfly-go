package parse

import (
	"errors"
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

func (p *Parser) predicate(parent *types.Expression) (*types.Expression, error) {
	left, err := p.equality(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isLogical) {
		_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

		if debug {
			fmt.Println("Parsing logical for lexeme:", p.previous().Text)
		}

		operator, err := types.NewOperator(p.previous().Text, false)

		if err != nil {
			return nil, err
		}

		right, err := p.equality(parent)

		if err != nil {
			return nil, err
		}

		exp := types.NewExpression(nil, operator, []*types.Expression{left, right})

		left = &exp
	}

	return left, nil
}

func (p *Parser) equality(parent *types.Expression) (*types.Expression, error) {
	left, err := p.comparison(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isEquality) {
		_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

		if debug {
			fmt.Println("Parsing equality for lexeme:", p.previous().Text)
		}

		operator, err := types.NewOperator(p.previous().Text, false)

		if err != nil {
			return nil, err
		}

		right, err := p.comparison(parent)

		if err != nil {
			return nil, err
		}

		exp := types.NewExpression(nil, operator, []*types.Expression{left, right})

		left = &exp
	}

	return left, nil
}

func (p *Parser) comparison(parent *types.Expression) (*types.Expression, error) {
	left, err := p.term(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isComparison) {
		_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

		if debug {
			fmt.Println("Parsing comparison for lexeme:", p.previous().Text)
		}

		operator, err := types.NewOperator(p.previous().Text, false)

		if err != nil {
			return nil, err
		}

		right, err := p.term(parent)

		if err != nil {
			return nil, err
		}

		exp := types.NewExpression(nil, operator, []*types.Expression{left, right})

		left = &exp
	}

	return left, nil
}

func (p *Parser) term(parent *types.Expression) (*types.Expression, error) {
	left, err := p.factor(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isTerm) {
		_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

		if debug {
			fmt.Println("Parsing term for lexeme:", p.previous().Text)
		}

		operator, err := types.NewOperator(p.previous().Text, false)

		if err != nil {
			return nil, err
		}

		right, err := p.factor(parent)

		if err != nil {
			return nil, err
		}

		exp := types.NewExpression(nil, operator, []*types.Expression{left, right})

		left = &exp
	}

	return left, nil
}

func (p *Parser) factor(parent *types.Expression) (*types.Expression, error) {
	left, err := p.unary(parent)

	if err != nil {
		return nil, err
	}

	for accept(p, isFactor) {
		_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

		if debug {
			fmt.Println("Parsing factor for lexeme:", p.previous().Text)
		}

		operator, err := types.NewOperator(p.previous().Text, false)

		if err != nil {
			return nil, err
		}

		right, err := p.unary(parent)

		if err != nil {
			return nil, err
		}

		exp := types.NewExpression(nil, operator, []*types.Expression{left, right})

		left = &exp
	}

	return left, nil
}

func (p *Parser) unary(parent *types.Expression) (*types.Expression, error) {
	if accept(p, isUnary) {
		_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

		if debug {
			fmt.Println("Parsing unary for lexeme:", p.previous().Text)
		}

		operator, err := types.NewOperator(p.previous().Text, false)

		if err != nil {
			return nil, err
		}

		right, err := p.unary(parent)

		if err != nil {
			return nil, err
		}

		exp := types.NewExpression(nil, operator, []*types.Expression{right})

		return &exp, nil
	}

	return p.atom(parent)
}

func (p *Parser) atom(parent *types.Expression) (*types.Expression, error) {
	if accept(p, isAtom) {
		_, debug := os.LookupEnv("GADFLY_DEBUG_PARSE")

		if debug {
			fmt.Println("Parsing atom for lexeme:", p.previous().Text)
		}

		operator, err := types.NewOperator(p.previous().Text, false)

		if err != nil {
			return nil, err
		}

		result := types.NewExpression(nil, operator, []*types.Expression{})

		return &result, nil
	}

	return nil, errors.New("expected expression but got <" + p.read().Text + ">")
}
