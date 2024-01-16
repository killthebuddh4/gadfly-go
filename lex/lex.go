package lex

func Lex(source string) ([]Lexeme, error) {
	lexer := NewLexer(source)

	for !lexer.isAtEnd() {
		lexer.Start = lexer.Current
		err := lexer.scan()

		if err != nil {
			return nil, err
		}
	}

	eof, err := NewEof(source)

	if err != nil {
		return nil, err
	}

	lexer.Tokens = append(lexer.Tokens, eof)

	return lexer.Tokens, nil
}
