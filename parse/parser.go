package parse

import (
	"errors"

	lib "github.com/killthebuddh4/gadflai/types"
)

type Parser struct {
	Lexemes []lib.Lexeme
	Current int
}

func accept(p *Parser, predicate func(lexeme lib.Lexeme) bool) bool {
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

func (p Parser) read() lib.Lexeme {
	return p.Lexemes[p.Current]
}

func (p Parser) previous() lib.Lexeme {
	return p.Lexemes[p.Current-1]
}

func (p Parser) isAtEnd() bool {
	return p.Current >= len(p.Lexemes)-1
}
