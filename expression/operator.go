package expression

import (
	"errors"
)

type Operator struct {
	Type  string
	Value string
}

var OPERATORS = []string{
	"program",
	"fn",
	"|",
	"def",
	"@",
	"val",
	"let",
	"if",
	"do",
	"and",
	"or",
	"while",
	"record",
	"merge",
	"delete",
	"keys",
	"values",
	"read",
	"write",
	"extract",
	"array",
	"get",
	"set",
	"for",
	"map",
	"filter",
	"reduce",
	"push",
	"pop",
	"shift",
	"unshift",
	"segment",
	"find",
	"splice",
	"reverse",
	"sort",
	"puts",
	"chars",
	"effect",
	"record",
	"+",
	"-",
	"*",
	"/",
	"!",
	"!=",
	"==",
	"<",
	"<=",
	">",
	">=",
	"&&",
	"||",
	"number",
	"string",
	"identifier",
	"true",
	"false",
	"nil",
}

func NewOperator(from string) (Operator, error) {
	for _, operator := range OPERATORS {
		if operator == from {
			return Operator{
				Type:  from,
				Value: from,
			}, nil
		}
	}

	switch string(from[0]) {
	case "\"":
		return Operator{
			Type:  "string",
			Value: from,
		}, nil
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return Operator{
			Type:  "number",
			Value: from,
		}, nil
	case "_", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "x", "w", "y", "z":
		return Operator{
			Type:  "identifier",
			Value: from,
		}, nil
	default:
		return Operator{}, errors.New("Could not resolve to operator from text: <" + from + ">")
	}
}
