package main

import (
	"errors"
	"fmt"
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
		_, err := p.expression(root)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) expression(parent *Expression) (Expression, error) {
	if p.accept([]string{"fn"}) {
		return p.lambda(parent)
	} else if p.accept(KEYWORDS) {
		return p.block(parent, p.previous().Type)
	} else if isDefined(parent, GetLexemeForToken(p.read())) {
		// isDefined doesn't accept the identifier, so we have to manually push it
		// along here
		if !p.accept([]string{"IDENTIFIER"}) {
			return Expression{}, errors.New("expected identifier")
		}
		fmt.Println("ADDING CALL")
		return p.block(parent, "call")
	} else {
		exp, err := p.logical()

		if err != nil {
			return Expression{}, err
		}

		return exp, nil
	}
}

func (p *Parser) block(parent *Expression, blockType string) (Expression, error) {
	operator := p.previous()

	operator.Type = blockType

	fmt.Println("Parsing block of type " + blockType)

	root := Expr(parent, operator)

	inputs := []Expression{}

	if blockType == "def" {
		fmt.Println("the next one is " + p.read().Type + " " + GetLexemeForToken(p.read()))

		input, err := p.logical()

		if err != nil {
			return Expression{}, err
		}

		inputs = append(inputs, input)

		identifier := p.previous()
		lexeme := GetLexemeForToken(identifier)
		parent.Definitions[lexeme] = nil
	}

	for !p.accept([]string{"end"}) && !p.isAtEnd() {

		fmt.Println("the next one is " + p.read().Type + " " + GetLexemeForToken(p.read()))
		input, err := p.expression(&root)

		if err != nil {
			return Expression{}, err
		}

		inputs = append(inputs, input)
	}

	root.Inputs = inputs

	fmt.Println("Finished parsing block of type " + blockType)
	fmt.Println("Terminating token was " + p.previous().Type + " " + GetLexemeForToken(p.previous()))
	fmt.Println("Next token is " + p.read().Type + " " + GetLexemeForToken(p.read()))

	parent.Inputs = append(parent.Inputs, root)
	return root, nil
}

func (p *Parser) lambda(parent *Expression) (Expression, error) {

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

	parent.Inputs = append(parent.Inputs, root)
	return root, nil
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

func isDefined(inExp *Expression, lexeme string) bool {
	if inExp == nil {
		return false
	}

	for kw, _ := range inExp.Definitions {
		if kw == lexeme {
			return true
		}
	}

	return isDefined(inExp.Parent, lexeme)
}

func getDefinition(inExp *Expression, lexeme string) (Lambda, error) {
	if inExp == nil {
		return nil, errors.New("symbol not found " + lexeme)
	}

	for kw, def := range inExp.Definitions {
		if kw == lexeme {
			return def, nil
		}
	}

	return getDefinition(inExp.Parent, lexeme)
}

func setDefinition(inExp *Expression, lexeme string, def Lambda) error {
	if inExp == nil {
		return errors.New("cannot set definition in nil expression")
	}

	inExp.Definitions[lexeme] = def

	return nil
}
