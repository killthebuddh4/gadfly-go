package lex

import "errors"

type Lexer struct {
	Source  string
	Tokens  []Lexeme
	Start   int
	Current int
}

func NewLexer(source string) Lexer {
	return Lexer{
		Source:  source,
		Tokens:  []Lexeme{},
		Start:   0,
		Current: 0,
	}
}

func (s *Lexer) advance() error {
	if s.isAtEnd() {
		return errors.New("unexpected end of file")
	}
	s.Current++

	return nil
}

func (s Lexer) readCurrent() (string, error) {
	if s.isAtEnd() {
		return "", errors.New("unexpected end of file")
	}
	return string(s.Source[s.Current]), nil
}

func (s Lexer) readNext() (string, error) {
	if s.Current+1 > len(s.Source) {
		return "", errors.New("unexpected end of file")
	}
	return string(s.Source[s.Current+1]), nil
}

func (s Lexer) isAtEnd() bool {
	return s.Current >= len(s.Source)
}
