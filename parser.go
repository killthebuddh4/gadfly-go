package main

import (
	"errors"
)

type Parser struct {
	Tokens  []Token
	Current int
}

func Parse(tokens []Token) ([]Expression, error) {
	parser := Parser{
		Tokens:  tokens,
		Current: 0,
	}

	return parser.program()
}

func (p *Parser) program() ([]Expression, error) {
	expressions := []Expression{}

	for !p.isAtEnd() {
		left, err := p.expression()

		if err != nil {
			return []Expression{}, err
		}

		expressions = append(expressions, left)
	}

	return expressions, nil
}

var EXPRESSIONS = []string{"def", "let", "call", "if", "get", "set", "do", "when", "then", "else", "and", "or", "array", "for", "map", "reduce", "filter"}

func (p *Parser) expression() (Expression, error) {
	if p.accept(EXPRESSIONS) {
		return p.block(p.previous().Type)
	} else if p.accept([]string{"fn"}) {
		return p.fn()
	} else {
		exp, err := p.logical()

		if err != nil {
			return Expression{}, err
		}

		return exp, nil
	}
}

func (p *Parser) block(blockType string) (Expression, error) {
	operator := p.previous()

	operator.Type = blockType

	expressions := []Expression{}

	for !p.accept([]string{"end"}) {
		expression, err := p.expression()

		if err != nil {
			return Expression{}, err
		}

		expressions = append(expressions, expression)
	}

	return Expression{
		Operator: operator,
		Inputs:   expressions,
	}, nil
}

func (p *Parser) fn() (Expression, error) {

	operator := p.previous()

	var parameters Expression

	if p.accept([]string{"PIPE"}) {
		pipe := p.previous()

		identifiers := []Expression{}

		for p.accept([]string{"IDENTIFIER"}) {
			identifiers = append(identifiers, Expression{
				Operator: p.previous(),
				Inputs:   []Expression{},
			})

			p.accept([]string{"COMMA"})
		}

		if !p.accept([]string{"PIPE"}) {
			return Expression{}, errors.New("expected closing pipe")
		}

		parameters = Expression{
			Operator: pipe,
			Inputs:   identifiers,
		}
	}

	block, err := p.block("do")

	if err != nil {
		return Expression{}, err
	}

	return Expression{
		Operator: operator,
		Inputs:   []Expression{parameters, block},
	}, nil
}

func (p *Parser) logical() (Expression, error) {
	left, err := p.equality()

	if err != nil {
		return Expression{}, err
	}

	for p.accept([]string{"&&", "||"}) {
		operator := p.previous()

		right, err := p.equality()

		if err != nil {
			return Expression{}, err
		}

		left = Expression{
			Operator: operator,
			Inputs:   []Expression{left, right},
		}
	}

	return left, nil
}

func (p *Parser) equality() (Expression, error) {
	left, err := p.comparison()

	if err != nil {
		return Expression{}, err
	}

	for p.accept([]string{"BANG_EQUAL", "EQUAL_EQUAL"}) {
		operator := p.previous()

		right, err := p.comparison()

		if err != nil {
			return Expression{}, err
		}

		left = Expression{
			Operator: operator,
			Inputs:   []Expression{left, right},
		}
	}

	return left, nil
}

func (p *Parser) comparison() (Expression, error) {
	left, err := p.term()

	if err != nil {
		return Expression{}, err
	}

	for p.accept([]string{"GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL"}) {
		operator := p.previous()

		right, err := p.term()

		if err != nil {
			return Expression{}, err
		}

		left = Expression{
			Operator: operator,
			Inputs:   []Expression{left, right},
		}
	}

	return left, nil
}

func (p *Parser) term() (Expression, error) {
	left, err := p.factor()

	if err != nil {
		return Expression{}, err
	}

	for p.accept([]string{"MINUS", "PLUS"}) {
		operator := p.previous()

		right, err := p.factor()

		if err != nil {
			return Expression{}, err
		}

		left = Expression{
			Operator: operator,
			Inputs:   []Expression{left, right},
		}
	}

	return left, nil
}

func (p *Parser) factor() (Expression, error) {
	left, err := p.unary()

	if err != nil {
		return Expression{}, err
	}

	for p.accept([]string{"SLASH", "STAR"}) {
		operator := p.previous()

		right, err := p.unary()

		if err != nil {
			return Expression{}, err
		}

		left = Expression{
			Operator: operator,
			Inputs:   []Expression{left, right},
		}
	}

	return left, nil
}

func (p *Parser) unary() (Expression, error) {

	if p.accept([]string{"BANG", "MINUS"}) {
		operator := p.previous()

		right, err := p.unary()

		if err != nil {
			return Expression{}, err
		}

		return Expression{
			Operator: operator,
			Inputs:   []Expression{right},
		}, nil
	}

	return p.atom()
}

func (p *Parser) atom() (Expression, error) {

	if p.accept([]string{"true", "false", "nil", "NUMBER", "STRING", "IDENTIFIER"}) {
		operator := p.previous()

		return Expression{
			Operator: operator,
			Inputs:   nil,
		}, nil
	}

	return Expression{}, errors.New("expected expression but got " + p.read().Type)
}

func (p *Parser) accept(tokenTypes []string) bool {
	for _, tokenType := range tokenTypes {
		if p.read().Type == tokenType {
			p.advance()
			return true
		}
	}

	return false
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
