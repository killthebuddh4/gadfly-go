package lex

import "github.com/killthebuddh4/gadflai/types"

func Lex(source string) ([]types.Lexeme, error) {
	lexer := NewLexer(source)

	for !lexer.isAtEnd() {
		lexer.Start = lexer.Current
		err := lexer.scan()

		if err != nil {
			return nil, err
		}
	}

	eof, err := types.NewEof(source)

	if err != nil {
		return nil, err
	}

	lexer.Tokens = append(lexer.Tokens, eof)

	return lexer.Tokens, nil
}
