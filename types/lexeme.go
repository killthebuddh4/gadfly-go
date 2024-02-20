package types

import "fmt"

type Lexeme struct {
	Start  int
	Length int
	Line   int
	Text   string
}

func NewLexeme(start int, length int, line int, lexeme string) (Lexeme, error) {
	return Lexeme{
		Start:  start,
		Length: length,
		Line:   line,
		Text:   lexeme,
	}, nil
}

func NewEof(forSource string) (Lexeme, error) {
	return NewLexeme(len(forSource), 0, 0, "EOF")
}

func PrintLexeme(lexeme Lexeme) {
	fmt.Printf("Lexeme: <%s> \n", lexeme.Text)
}
