package main

import "fmt"

type Token struct {
	Type   string
	Start  int
	Length int
}

func GetLexemeForToken(token Token) string {
	return GetSource()[token.Start : token.Start+token.Length]
}

func PrintToken(token Token) {
	lexeme := GetLexemeForToken(token)
	fmt.Printf("Type: %s, Lexeme: <%s> \n", token.Type, lexeme)
}
