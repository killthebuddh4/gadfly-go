package main

import (
	"errors"
	"fmt"
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

	fmt.Println("Parsing")

	return parser.program()
}

func (p *Parser) program() ([]Expression, error) {
	fmt.Println("Parsing program")

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

func (p *Parser) expression() (Expression, error) {
	fmt.Println("Parsing expression")

	if p.accept([]string{"let"}) {
		return p.declaration()
	} else if p.accept([]string{"do"}) {
		return p.block()
	} else {
		exp, err := p.equality()

		if err != nil {
			return Expression{}, err
		}

		if !p.accept([]string{"SEMICOLON"}) {
			return Expression{}, errors.New("expected semicolon after expression")
		}

		return exp, nil
	}
}

func (p *Parser) declaration() (Expression, error) {
	fmt.Println("Parsing declaration")

	operator := p.previous()

	if !p.accept([]string{"IDENTIFIER"}) {
		return Expression{}, errors.New("expected identifier after let")
	}

	identifier := p.previous()

	var value Expression

	if p.accept([]string{"do"}) {
		v, err := p.block()

		if err != nil {
			return Expression{}, err
		}

		value = v
	} else {
		v, err := p.equality()

		if err != nil {
			return Expression{}, err
		}

		if !p.accept([]string{"SEMICOLON"}) {
			return Expression{}, errors.New("expected semicolon after declaration")
		}

		value = v
	}

	return Expression{
		Operator: operator,
		Inputs: []Expression{{
			Operator: identifier,
			Inputs:   []Expression{},
		}, value},
	}, nil
}

func (p *Parser) block() (Expression, error) {
	fmt.Println("Parsing block")

	operator := p.previous()

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

func (p *Parser) equality() (Expression, error) {
	fmt.Println("Parsing equality")

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
	fmt.Println("Parsing comparison")

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
	fmt.Println("Parsing term")

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
	fmt.Println("Parsing factor")

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
	fmt.Println("Parsing unary")

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
	fmt.Println("Parsing atom")

	if p.accept([]string{"TRUE", "FALSE", "NIL", "NUMBER", "STRING", "IDENTIFIER"}) {
		operator := p.previous()

		fmt.Println("Parsing literal of type ", operator.Type)

		return Expression{
			Operator: operator,
			Inputs:   nil,
		}, nil
	}

	if p.accept([]string{"LEFT_PAREN"}) {
		expr, err := p.expression()

		if err != nil {
			return Expression{}, err
		}

		if !p.accept([]string{"RIGHT_PAREN"}) {
			return Expression{}, errors.New("expected right paren")
		}

		return expr, nil
	}

	fmt.Println("Peek", p.read().Type)

	return Expression{}, errors.New("expected expression")
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
