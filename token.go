package main

import (
	"fmt"
)

type Token struct {
	Type   string
	Start  int
	Length int
	Lexeme string
}

type Tokens struct {
	Fn               string
	Pipe             string
	Def              string
	Call             string
	Val              string
	Edit             string
	If               string
	Do               string
	And              string
	Or               string
	Array            string
	Get              string
	Set              string
	For              string
	Map              string
	Reduce           string
	Filter           string
	Push             string
	Pop              string
	End              string
	Effect           string
	Plus             string
	Minus            string
	Multiply         string
	Divide           string
	Bang             string
	BangEqual        string
	EqualEqual       string
	LessThan         string
	LessThanEqual    string
	GreaterThan      string
	GreaterThanEqual string
	Conjunction      string
	Disjunction      string
	Number           string
	String           string
	Identifier       string
	True             string
	False            string
	Nil              string
	EndOfFile        string
}

var TOKENS = Tokens{
	Fn:               "fn",
	Pipe:             "|",
	Def:              "def",
	Call:             "call",
	Val:              "val",
	Edit:             "let",
	If:               "if",
	Do:               "do",
	And:              "and",
	Or:               "or",
	Array:            "array",
	Get:              "get",
	Set:              "set",
	For:              "for",
	Map:              "map",
	Reduce:           "reduce",
	Filter:           "filter",
	Push:             "push",
	Pop:              "pop",
	End:              "end",
	Effect:           "effect",
	Plus:             "+",
	Minus:            "-",
	Multiply:         "*",
	Divide:           "/",
	Bang:             "!",
	BangEqual:        "!=",
	EqualEqual:       "==",
	LessThan:         "<",
	LessThanEqual:    "<=",
	GreaterThan:      ">",
	GreaterThanEqual: ">=",
	Conjunction:      "&&",
	Disjunction:      "||",
	Number:           "number",
	String:           "string",
	Identifier:       "identifier",
	True:             "true",
	False:            "false",
	Nil:              "nil",
	EndOfFile:        "EOF",
}

func isLogical(token Token) bool {
	switch token.Type {
	case TOKENS.Conjunction, TOKENS.Disjunction:
		return true
	default:
		return false
	}
}

func isEquality(token Token) bool {
	switch token.Type {
	case TOKENS.BangEqual, TOKENS.EqualEqual:
		return true
	default:
		return false
	}
}

func isComparison(token Token) bool {
	switch token.Type {
	case TOKENS.BangEqual, TOKENS.LessThan, TOKENS.LessThanEqual, TOKENS.GreaterThan, TOKENS.GreaterThanEqual:
		return true
	default:
		return false
	}
}

func isTerm(token Token) bool {
	switch token.Type {
	case TOKENS.Plus, TOKENS.Minus:
		return true
	default:
		return false
	}
}

func isFactor(token Token) bool {
	switch token.Type {
	case TOKENS.Multiply, TOKENS.Divide:
		return true
	default:
		return false
	}
}

func isUnary(token Token) bool {
	switch token.Type {
	case TOKENS.Minus, TOKENS.Bang:
		return true
	default:
		return false
	}
}

func isAtom(token Token) bool {
	switch token.Type {
	case TOKENS.Number, TOKENS.String, TOKENS.Identifier, TOKENS.True, TOKENS.False, TOKENS.Nil:
		return true
	default:
		return false
	}
}

func isBlockStart(token Token) bool {
	switch token.Type {
	case TOKENS.Def, TOKENS.Call, TOKENS.Val, TOKENS.Edit, TOKENS.If, TOKENS.Do, TOKENS.And, TOKENS.Or, TOKENS.Array, TOKENS.Get, TOKENS.Set, TOKENS.For, TOKENS.Map, TOKENS.Reduce, TOKENS.Filter, TOKENS.Push, TOKENS.Pop, TOKENS.Effect:
		return true
	default:
		return false
	}
}

func isPipe(token Token) bool {
	switch token.Type {
	case TOKENS.Pipe:
		return true
	default:
		return false
	}
}

func isIdentifier(token Token) bool {
	switch token.Type {
	case TOKENS.Identifier:
		return true
	default:
		return false
	}
}

func isFn(token Token) bool {
	switch token.Type {
	case TOKENS.Fn:
		return true
	default:
		return false
	}
}

func isEnd(token Token) bool {
	switch token.Type {
	case TOKENS.End:
		return true
	default:
		return false
	}
}

func isTokenType(tokenType string) bool {
	switch tokenType {
	case TOKENS.Fn, TOKENS.Pipe, TOKENS.Def, TOKENS.Call, TOKENS.Val, TOKENS.Edit, TOKENS.If, TOKENS.Do, TOKENS.And, TOKENS.Or, TOKENS.Array, TOKENS.Get, TOKENS.Set, TOKENS.For, TOKENS.Map, TOKENS.Reduce, TOKENS.Filter, TOKENS.Push, TOKENS.Pop, TOKENS.End, TOKENS.Effect, TOKENS.Plus, TOKENS.Minus, TOKENS.Multiply, TOKENS.Divide, TOKENS.Bang, TOKENS.BangEqual, TOKENS.EqualEqual, TOKENS.LessThan, TOKENS.LessThanEqual, TOKENS.GreaterThan, TOKENS.GreaterThanEqual, TOKENS.Number, TOKENS.String, TOKENS.Identifier, TOKENS.True, TOKENS.False, TOKENS.Nil, TOKENS.EndOfFile:
		return true
	default:
		return false
	}
}

func GetLexemeForToken(token Token) string {
	return GetSource()[token.Start : token.Start+token.Length]
}

func PrintToken(token Token) {
	lexeme := GetLexemeForToken(token)
	fmt.Printf("Type: %s, Lexeme: <%s> \n", token.Type, lexeme)
}

func NewToken(tokenType string, start int, length int, lexeme string) (Token, error) {
	if !isTokenType(tokenType) {
		return Token{}, fmt.Errorf("invalid token type: %s", tokenType)
	}

	return Token{
		Type:   tokenType,
		Start:  start,
		Length: length,
		Lexeme: lexeme,
	}, nil
}

func EofToken(forSource string) (Token, error) {
	return NewToken(TOKENS.EndOfFile, len(forSource), 0, "EOF")
}
