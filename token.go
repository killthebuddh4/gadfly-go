package main

import "fmt"

type Token struct {
	Type   string
	Start  int
	Length int
}

func GetLexemeForToken(source string, token Token) string {
	return source[token.Start : token.Start+token.Length]
}

func PrintToken(source string, token Token) {
	lexeme := GetLexemeForToken(source, token)
	fmt.Printf("Type: %s, Lexeme: <%s> \n", token.Type, lexeme)
}
