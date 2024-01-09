package main

import "errors"

type Blocks struct {
	Fn     string
	Def    string
	Call   string
	Val    string
	Edit   string
	If     string
	Do     string
	And    string
	Or     string
	Array  string
	Get    string
	Set    string
	For    string
	Map    string
	Reduce string
	Filter string
	Push   string
	Pop    string
	End    string
	Effect string
}

var BLOCKS = Blocks{
	Fn:     "fn",
	Def:    "def",
	Call:   "call",
	Val:    "val",
	Edit:   "let",
	If:     "if",
	Do:     "do",
	And:    "and",
	Or:     "or",
	Array:  "array",
	Get:    "get",
	Set:    "set",
	For:    "for",
	Map:    "map",
	Reduce: "reduce",
	Filter: "filter",
	Push:   "push",
	Pop:    "pop",
	End:    "end",
	Effect: "effect",
}

func GetBlock(lexeme string) (string, error) {
	switch lexeme {
	case BLOCKS.Fn:
		return BLOCKS.Fn, nil
	case BLOCKS.Def:
		return BLOCKS.Def, nil
	case BLOCKS.Call:
		return BLOCKS.Call, nil
	case BLOCKS.Val:
		return BLOCKS.Val, nil
	case BLOCKS.Edit:
		return BLOCKS.Edit, nil
	case BLOCKS.If:
		return BLOCKS.If, nil
	case BLOCKS.Do:
		return BLOCKS.Do, nil
	case BLOCKS.And:
		return BLOCKS.And, nil
	case BLOCKS.Or:
		return BLOCKS.Or, nil
	case BLOCKS.Array:
		return BLOCKS.Array, nil
	case BLOCKS.Get:
		return BLOCKS.Get, nil
	case BLOCKS.Set:
		return BLOCKS.Set, nil
	case BLOCKS.For:
		return BLOCKS.For, nil
	case BLOCKS.Map:
		return BLOCKS.Map, nil
	case BLOCKS.Reduce:
		return BLOCKS.Reduce, nil
	case BLOCKS.Filter:
		return BLOCKS.Filter, nil
	case BLOCKS.Push:
		return BLOCKS.Push, nil
	case BLOCKS.Pop:
		return BLOCKS.Pop, nil
	case BLOCKS.End:
		return BLOCKS.End, nil
	case BLOCKS.Effect:
		return BLOCKS.Effect, nil
	default:
		return "", errors.New("unexpected block")
	}
}

type Special struct {
	Pipe rune
}

var SPECIAL = Special{
	Pipe: '|',
}

type Comment struct {
	Line rune
}

var COMMENT = Comment{
	Line: '#',
}

type Operators struct {
	Plus        rune
	Minus       rune
	Multiply    rune
	Divide      rune
	Bang        rune
	Equal       rune
	LessThan    rune
	GreaterThan rune
	Ampersand   rune
	Pipe        rune
}

var OPERATORS = Operators{
	Plus:        '+',
	Minus:       '-',
	Multiply:    '*',
	Divide:      '/',
	Equal:       '=',
	Bang:        '!',
	LessThan:    '<',
	GreaterThan: '>',
	Ampersand:   '&',
	Pipe:        '|',
}

type WhiteSpace struct {
	NewLine rune
	Space   rune
	Tab     rune
	Return  rune
}

var WHITESPACE = WhiteSpace{
	NewLine: '\n',
	Space:   ' ',
	Tab:     '\t',
	Return:  '\r',
}

type Constants struct {
	True  string
	False string
	Nil   string
}

var CONSTANTS = Constants{
	True:  "true",
	False: "false",
	Nil:   "nil",
}

type Numbers struct {
	Dot   rune
	Zero  rune
	One   rune
	Two   rune
	Three rune
	Four  rune
	Five  rune
	Six   rune
	Seven rune
	Eight rune
	Nine  rune
}

var NUMBERS = Numbers{
	Dot:   '.',
	Zero:  '0',
	One:   '1',
	Two:   '2',
	Three: '3',
	Four:  '4',
	Five:  '5',
	Six:   '6',
	Seven: '7',
	Eight: '8',
	Nine:  '9',
}

type Strings struct {
	Quote rune
}

var STRINGS = Strings{
	Quote: '"',
}

type Identifiers struct {
	Underscore rune
	LowerA     rune
	LowerB     rune
	LowerC     rune
	LowerD     rune
	LowerE     rune
	LowerF     rune
	LowerG     rune
	LowerH     rune
	LowerI     rune
	LowerJ     rune
	LowerK     rune
	LowerL     rune
	LowerM     rune
	LowerN     rune
	LowerO     rune
	LowerP     rune
	LowerQ     rune
	LowerR     rune
	LowerS     rune
	LowerT     rune
	LowerU     rune
	LowerV     rune
	LowerW     rune
	LowerX     rune
	LowerY     rune
	LowerZ     rune
	UpperA     rune
	UpperB     rune
	UpperC     rune
	UpperD     rune
	UpperE     rune
	UpperF     rune
	UpperG     rune
	UpperH     rune
	UpperI     rune
	UpperJ     rune
	UpperK     rune
	UpperL     rune
	UpperM     rune
	UpperN     rune
	UpperO     rune
	UpperP     rune
	UpperQ     rune
	UpperR     rune
	UpperS     rune
	UpperT     rune
	UpperU     rune
	UpperV     rune
	UpperW     rune
	UpperX     rune
	UpperY     rune
	UpperZ     rune
}

var IDENTIFIERS = Identifiers{
	Underscore: '_',
	LowerA:     'a',
	LowerB:     'b',
	LowerC:     'c',
	LowerD:     'd',
	LowerE:     'e',
	LowerF:     'f',
	LowerG:     'g',
	LowerH:     'h',
	LowerI:     'i',
	LowerJ:     'j',
	LowerK:     'k',
	LowerL:     'l',
	LowerM:     'm',
	LowerN:     'n',
	LowerO:     'o',
	LowerP:     'p',
	LowerQ:     'q',
	LowerR:     'r',
	LowerS:     's',
	LowerT:     't',
	LowerU:     'u',
	LowerV:     'v',
	LowerW:     'w',
	LowerX:     'x',
	LowerY:     'y',
	LowerZ:     'z',
	UpperA:     'A',
	UpperB:     'B',
	UpperC:     'C',
	UpperD:     'D',
	UpperE:     'E',
	UpperF:     'F',
	UpperG:     'G',
	UpperH:     'H',
	UpperI:     'I',
	UpperJ:     'J',
	UpperK:     'K',
	UpperL:     'L',
	UpperM:     'M',
	UpperN:     'N',
	UpperO:     'O',
	UpperP:     'P',
	UpperQ:     'Q',
	UpperR:     'R',
	UpperS:     'S',
	UpperT:     'T',
	UpperU:     'U',
	UpperV:     'V',
	UpperW:     'W',
	UpperX:     'X',
	UpperY:     'Y',
	UpperZ:     'Z',
}

func IsIdentifierChar(c rune) bool {
	switch c {
	case IDENTIFIERS.Underscore, IDENTIFIERS.LowerA, IDENTIFIERS.LowerB, IDENTIFIERS.LowerC, IDENTIFIERS.LowerD, IDENTIFIERS.LowerE, IDENTIFIERS.LowerF, IDENTIFIERS.LowerG, IDENTIFIERS.LowerH, IDENTIFIERS.LowerI, IDENTIFIERS.LowerJ, IDENTIFIERS.LowerK, IDENTIFIERS.LowerL, IDENTIFIERS.LowerM, IDENTIFIERS.LowerN, IDENTIFIERS.LowerO, IDENTIFIERS.LowerP, IDENTIFIERS.LowerQ, IDENTIFIERS.LowerR, IDENTIFIERS.LowerS, IDENTIFIERS.LowerT, IDENTIFIERS.LowerU, IDENTIFIERS.LowerV, IDENTIFIERS.LowerW, IDENTIFIERS.LowerX, IDENTIFIERS.LowerY, IDENTIFIERS.LowerZ, IDENTIFIERS.UpperA, IDENTIFIERS.UpperB, IDENTIFIERS.UpperC, IDENTIFIERS.UpperD, IDENTIFIERS.UpperE, IDENTIFIERS.UpperF, IDENTIFIERS.UpperG, IDENTIFIERS.UpperH, IDENTIFIERS.UpperI, IDENTIFIERS.UpperJ, IDENTIFIERS.UpperK, IDENTIFIERS.UpperL, IDENTIFIERS.UpperM, IDENTIFIERS.UpperN, IDENTIFIERS.UpperO, IDENTIFIERS.UpperP, IDENTIFIERS.UpperQ, IDENTIFIERS.UpperR, IDENTIFIERS.UpperS, IDENTIFIERS.UpperT, IDENTIFIERS.UpperU, IDENTIFIERS.UpperV, IDENTIFIERS.UpperW, IDENTIFIERS.UpperX, IDENTIFIERS.UpperY, IDENTIFIERS.UpperZ:
		return true
	default:
		return false
	}
}
