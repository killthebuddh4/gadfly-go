package lex

import "fmt"

type Lexeme struct {
	Type   string
	Start  int
	Length int
	Text   string
}

func NewLexeme(kind string, start int, length int, lexeme string) (Lexeme, error) {
	return Lexeme{
		Type:   kind,
		Start:  start,
		Length: length,
		Text:   lexeme,
	}, nil
}

func NewEof(forSource string) (Lexeme, error) {
	return NewLexeme(KEYWORDS.EndOfFile, len(forSource), 0, "EOF")
}

func PrintLexeme(lexeme Lexeme) {
	fmt.Printf("Type: %s, Lexeme: <%s> \n", lexeme.Type, lexeme.Text)
}
