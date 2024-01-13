package lex

var source *string

func GetSource() string {
	return *source
}

func SetSource(s string) {
	source = &s
}

var tokens *[]Lexeme

func GetTokens() []Lexeme {
	return *tokens
}

func SetTokens(t []Lexeme) {
	tokens = &t
}
