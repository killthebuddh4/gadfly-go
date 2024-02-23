package parse

import (
	"errors"

	lib "github.com/killthebuddh4/gadflai/types"
)

type Parser struct {
	Lexemes []lib.Lexeme
	Current int
}

type Predicate func(lexeme lib.Lexeme) bool

func accept(p *Parser, predicates ...Predicate) bool {
	if len(predicates) == 0 {
		// TODO accept should return an error.
		panic("accept called with no predicates")
	}

	token := p.read()

	for _, predicate := range predicates {
		if predicate(token) {
			p.advance()
			return true
		}
	}

	return false
}

func acc(p *Parser, cur string, words []string) bool {
	if len(words) == 0 {
		// TODO accept should return an error.
		panic("accept called with no words")
	}

	for _, word := range words {
		if cur == word {
			p.advance()
			return true
		}
	}

	return false
}

func couldAccept(p *Parser, cur string, words []string) bool {
	if len(words) == 0 {
		// TODO accept should return an error.
		panic("accept called with no words")
	}

	for _, word := range words {
		if cur == word {
			return true
		}
	}

	return false
}

func (p *Parser) backup() error {
	if p.Current == 0 {
		return errors.New("cannot backup")
	}

	p.Current--

	return nil
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
