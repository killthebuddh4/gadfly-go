package main

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func Expr(parent *Expression, variant string, operator Token) *Expression {
	id, err := uuid.NewRandom()

	if err != nil {
		panic(err)
	}

	expr := Expression{
		Id:           id.String(),
		Parent:       parent,
		Operator:     operator,
		Variant:      variant,
		Children:     []*Expression{},
		Parameters:   []string{},
		Trajectories: []*Trajectory{},
		Definitions:  []Definition{},
	}

	if parent != nil {
		parent.Children = append(parent.Children, &expr)
	}

	return &expr
}

func RootExpr() Expression {
	return Expression{
		Parent: nil,
		Operator: Token{
			Type:   TOKENS.Root,
			Lexeme: "ROOT",
			Start:  -1,
			Length: -1,
		},
		Variant:      VARIANTS.Root,
		Children:     []*Expression{},
		Parameters:   []string{},
		Trajectories: []*Trajectory{},
		Definitions:  BUILTINS,
	}
}

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

func (p *Parser) backup() error {
	if p.isAtEnd() {
		return errors.New("unexpected end of file")
	}

	p.Current--

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

func PrintExp(expr *Expression) {
	fmt.Println("Printing expression tree :)")
	printExp(expr, 0)
}

func printExp(expr *Expression, indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}

	fmt.Println("<", expr.Operator.Lexeme, ">")

	for _, child := range expr.Children {
		printExp(child, indent+4)
	}
}
