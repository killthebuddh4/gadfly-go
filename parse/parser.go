package parse

import (
	"errors"

	"github.com/killthebuddh4/gadflai/lex"
)

type Parser struct {
	Lexemes []lex.Lexeme
	Current int
}

func accept(p *Parser, predicate func(lexeme lex.Lexeme) bool) bool {
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

func (p Parser) read() lex.Lexeme {
	return p.Lexemes[p.Current]
}

func (p Parser) previous() lex.Lexeme {
	return p.Lexemes[p.Current-1]
}

func (p Parser) isAtEnd() bool {
	return p.Current >= len(p.Lexemes)-1
}
