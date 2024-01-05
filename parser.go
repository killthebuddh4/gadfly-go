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

	return parser.program()
}

func (p *Parser) program() ([]Expression, error) {
	expressions := []Expression{}

	for !p.isAtEnd() {
		left, err := p.sexp()

		if err != nil {
			return []Expression{}, err
		}

		expressions = append(expressions, left)
	}

	return expressions, nil
}

var EXPRESSIONS = []string{"def", "if", "set", "do", "when", "then", "else", "and", "or", "array", "for", "map", "reduce", "filter"}

func (p *Parser) sexp() (Expression, error) {
	if p.accept(EXPRESSIONS) {
		fmt.Println("Block is of type", p.previous().Type)
		return p.sblock()
	} else if p.accept([]string{"fn"}) {
		return p.sfn()
	} else {
		exp, err := p.logical()

		if err != nil {
			return Expression{}, err
		}

		return exp, nil
	}
}

func (p *Parser) sblock() (Expression, error) {
	operator := p.previous()

	fmt.Println("Parsing sblock of type", operator.Type)

	expressions := []Expression{}

	for !p.accept([]string{"end"}) {
		fmt.Println("looking at ", p.read().Type, GetLexemeForToken(p.read()))
		expression, err := p.sexp()

		if err != nil {
			return Expression{}, err
		}

		expressions = append(expressions, expression)
	}

	fmt.Println("Done parsing sblock of type", operator.Type)

	return Expression{
		Operator: operator,
		Inputs:   expressions,
	}, nil
}

func (p *Parser) sfn() (Expression, error) {
	fmt.Println("Parsing sfn")

	operator := p.previous()

	var parameters Expression

	if p.accept([]string{"PIPE"}) {
		fmt.Println("Parsing PIPE")
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

		fmt.Println("Done parsing PIPE")
	}

	if !p.accept([]string{"do"}) {
		return Expression{}, errors.New("expected do after parameters")
	}

	block, err := p.sblock()

	if err != nil {
		return Expression{}, err
	}

	return Expression{
		Operator: operator,
		Inputs:   []Expression{parameters, block},
	}, nil
}

func (p *Parser) expression() (Expression, error) {
	if p.accept([]string{"def"}) {
		return p.declaration()
	} else if p.accept([]string{"fn"}) {
		return p.fun()
	} else if p.accept([]string{"array"}) {
		return p.array()
	} else if p.accept([]string{"set"}) {
		return p.set()
	} else if p.accept([]string{"do"}) {
		return p.block()
	} else if p.accept([]string{"if"}) {
		return p.parseIf()
	} else {
		exp, err := p.logical()

		if err != nil {
			return Expression{}, err
		}

		return exp, nil
	}
}

func (p *Parser) array() (Expression, error) {
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

func (p *Parser) parseIf() (Expression, error) {
	operator := p.previous()

	condition, err := p.logical()

	if err != nil {
		return Expression{}, err
	}

	if !p.accept([]string{"then"}) {
		return Expression{}, errors.New("expected then after condition")
	}

	thenExp, err := p.block()

	if err != nil {
		return Expression{}, err
	}

	elseExp, err := p.block()

	if err != nil {
		return Expression{}, err
	}

	return Expression{
		Operator: operator,
		Inputs:   []Expression{condition, thenExp, elseExp},
	}, nil
}

func (p *Parser) declaration() (Expression, error) {
	operator := p.previous()

	if !p.accept([]string{"IDENTIFIER"}) {
		return Expression{}, errors.New("expected identifier after def")
	}

	identifier := p.previous()

	value, err := p.expression()

	if err != nil {
		return Expression{}, err
	}

	return Expression{
		Operator: operator,
		Inputs: []Expression{{
			Operator: identifier,
			Inputs:   []Expression{},
		}, value},
	}, nil
}

func (p *Parser) set() (Expression, error) {
	operator := p.previous()

	if !p.accept([]string{"IDENTIFIER"}) {
		return Expression{}, errors.New("expected identifier after def")
	}

	identifier := p.previous()

	value, err := p.expression()

	if err != nil {
		return Expression{}, err
	}

	return Expression{
		Operator: operator,
		Inputs: []Expression{{
			Operator: identifier,
			Inputs:   []Expression{},
		}, value},
	}, nil
}

func (p *Parser) fun() (Expression, error) {
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

	block, err := p.block()

	if err != nil {
		return Expression{}, err
	}

	return Expression{
		Operator: operator,
		Inputs:   []Expression{parameters, block},
	}, nil
}

func (p *Parser) block() (Expression, error) {
	operator := p.previous()

	// TOOD
	operator.Type = "do"

	expressions := []Expression{}

	for !p.accept([]string{"end", "else"}) {
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

	return p.call()
}

func (p *Parser) call() (Expression, error) {
	left, err := p.atom()

	if err != nil {
		return Expression{}, err
	}

	for p.accept([]string{"LEFT_PAREN"}) {
		operator := p.previous()

		arguments := []Expression{}

		for !p.accept([]string{"RIGHT_PAREN"}) {
			argument, err := p.expression()

			if err != nil {
				return Expression{}, err
			}

			arguments = append(arguments, argument)

			p.accept([]string{"COMMA"})
		}

		left = Expression{
			Operator: operator,
			Inputs:   append([]Expression{left}, arguments...),
		}
	}

	return left, nil
}

func (p *Parser) atom() (Expression, error) {
	if p.accept([]string{"true", "false", "nil", "NUMBER", "STRING", "IDENTIFIER"}) {
		operator := p.previous()

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
