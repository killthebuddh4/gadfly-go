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
		left, err := p.expression(nil)

		if err != nil {
			return []Expression{}, err
		}

		expressions = append(expressions, left)
	}

	return expressions, nil
}

var EXPRESSIONS = []string{"def", "val", "let", "call", "if", "get", "set", "do", "when", "then", "else", "and", "or", "array", "for", "map", "reduce", "filter"}

func (p *Parser) expression(parent *Expression) (Expression, error) {
	names := []string{}
	if parent != nil {
		for name := range parent.Keywords {
			names = append(names, name)
		}
	}

	if p.accept(EXPRESSIONS) {
		return p.block(parent, p.previous().Type)
	} else if p.defined(names) {
		fmt.Println("PREV TYPE", p.previous().Type)
		return p.block(parent, p.previous().Type)
	} else if p.accept([]string{"fn"}) {
		return p.fn(parent)
	} else {
		exp, err := p.logical(parent)

		if err != nil {
			return Expression{}, err
		}

		return exp, nil
	}
}

func (p *Parser) block(parent *Expression, blockType string) (Expression, error) {
	operator := p.previous()

	operator.Type = blockType

	root := Expr(parent, operator)

	inputs := []Expression{}

	for !p.accept([]string{"end"}) {
		input, err := p.expression(&root)

		if err != nil {
			return Expression{}, err
		}

		inputs = append(inputs, input)
	}

	root.Inputs = inputs

	return root, nil
}

func (p *Parser) fn(parent *Expression) (Expression, error) {

	operator := p.previous()

	root := Expr(parent, operator)

	var parameters Expression

	if p.accept([]string{"PIPE"}) {
		pipe := p.previous()

		identifiers := []Expression{}

		for p.accept([]string{"IDENTIFIER"}) {
			identifiers = append(identifiers, Expr(&root, p.previous()))

			p.accept([]string{"COMMA"})
		}

		if !p.accept([]string{"PIPE"}) {
			return Expression{}, errors.New("expected closing pipe")
		}

		parameters = Expr(&root, pipe)
		parameters.Inputs = identifiers
	}

	block, err := p.block(&root, "do")

	if err != nil {
		return Expression{}, err
	}

	root.Inputs = []Expression{parameters, block}

	return root, nil
}

func (p *Parser) logical(parent *Expression) (Expression, error) {
	left, err := p.equality(parent)

	if err != nil {
		return Expression{}, err
	}

	for p.accept([]string{"&&", "||"}) {
		operator := p.previous()

		right, err := p.equality(parent)

		if err != nil {
			return Expression{}, err
		}

		left = Expr(parent, operator)
		left.Inputs = []Expression{left, right}
	}

	return left, nil
}

func (p *Parser) equality(parent *Expression) (Expression, error) {
	left, err := p.comparison(parent)

	if err != nil {
		return Expression{}, err
	}

	for p.accept([]string{"BANG_EQUAL", "EQUAL_EQUAL"}) {
		operator := p.previous()

		right, err := p.comparison(parent)

		if err != nil {
			return Expression{}, err
		}

		left = Expr(parent, operator)
		left.Inputs = []Expression{left, right}
	}

	return left, nil
}

func (p *Parser) comparison(parent *Expression) (Expression, error) {
	left, err := p.term(parent)

	if err != nil {
		return Expression{}, err
	}

	for p.accept([]string{"GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL"}) {
		operator := p.previous()

		right, err := p.term(parent)

		if err != nil {
			return Expression{}, err
		}

		left = Expr(parent, operator)
		left.Inputs = []Expression{left, right}
	}

	return left, nil
}

func (p *Parser) term(parent *Expression) (Expression, error) {
	left, err := p.factor(parent)

	if err != nil {
		return Expression{}, err
	}

	for p.accept([]string{"MINUS", "PLUS"}) {
		operator := p.previous()

		right, err := p.factor(parent)

		if err != nil {
			return Expression{}, err
		}

		left = Expr(parent, operator)
		left.Inputs = []Expression{left, right}
	}

	return left, nil
}

func (p *Parser) factor(parent *Expression) (Expression, error) {
	left, err := p.unary(parent)

	if err != nil {
		return Expression{}, err
	}

	for p.accept([]string{"SLASH", "STAR"}) {
		operator := p.previous()

		right, err := p.unary(parent)

		if err != nil {
			return Expression{}, err
		}

		left = Expr(parent, operator)
		left.Inputs = []Expression{left, right}
	}

	return left, nil
}

func (p *Parser) unary(parent *Expression) (Expression, error) {

	if p.accept([]string{"BANG", "MINUS"}) {
		operator := p.previous()

		right, err := p.unary(parent)

		if err != nil {
			return Expression{}, err
		}

		e := Expr(parent, operator)
		e.Inputs = []Expression{right}
		return e, nil
	}

	return p.atom(parent)
}

func (p *Parser) atom(parent *Expression) (Expression, error) {

	if p.accept([]string{"true", "false", "nil", "NUMBER", "STRING", "IDENTIFIER"}) {
		operator := p.previous()

		return Expr(parent, operator), nil
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

func (p *Parser) defined(names []string) bool {
	for _, name := range names {
		token := p.read()
		lexeme := GetLexemeForToken(token)
		if lexeme == name {
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
