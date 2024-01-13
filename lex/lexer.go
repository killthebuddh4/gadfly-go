package lex

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
