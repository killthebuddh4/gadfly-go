package main

var source *string

func GetSource() string {
	return *source
}

func SetSource(s string) {
	source = &s
}

var tokens *[]Token

func GetTokens() []Token {
	return *tokens
}

func SetTokens(t []Token) {
	tokens = &t
}

var expression *Expression

func GetExpression() Expression {
	return *expression
}

func SetExpression(e Expression) {
	expression = &e
}
