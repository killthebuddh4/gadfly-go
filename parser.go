package main

import (
	"errors"
	"fmt"
)

type Expression struct {
	Operator Token
	Inputs   []Expression
}

type Parser struct {
	Tokens  []Token
	Current int
}

func (p *Parser) Parse() (Expression, error) {
	fmt.Println("Parsing")

	return p.expression()
}

func (p *Parser) expression() (Expression, error) {
	fmt.Println("Parsing expression")

	return p.equality()
}

func (p *Parser) equality() (Expression, error) {
	fmt.Println("Parsing equality")

	left, err := p.comparison()

	if err != nil {
		return Expression{}, err
	}

	for p.accept([]string{"bang_equal", "equal_equal"}) {
		operator := p.read()

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
		operator := p.read()

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
		operator := p.read()

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
		operator := p.read()

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
		operator := p.read()

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

	if p.accept([]string{"TRUE", "FALSE", "NIL", "NUMBER", "STRING"}) {
		operator := p.read()

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

		// TODO, The book uses a grouping expression here, but I don't know how
		// we'll use it.
		return expr, nil
	}

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

func (p Parser) isAtEnd() bool {
	return p.Current >= len(p.Tokens)
}

func (exp Expression) print() {
	str := fmt.Sprintf("Operator: %s ", exp.Operator.Type)
	for _, input := range exp.Inputs {
		str += fmt.Sprintf("Input: %s ", input.Operator.Type)
	}

	fmt.Println(str)

	for _, input := range exp.Inputs {
		input.print()
	}
}
