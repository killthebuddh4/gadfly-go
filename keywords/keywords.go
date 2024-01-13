package keywords

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
	While  string
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
	Call:   "@",
	Val:    "val",
	Edit:   "let",
	If:     "if",
	Do:     "do",
	And:    "and",
	Or:     "or",
	While:  "while",
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

func IsBlock(text string) bool {
	switch text {
	case BLOCKS.Def, BLOCKS.Call, BLOCKS.Val, BLOCKS.Edit, BLOCKS.If, BLOCKS.Do, BLOCKS.And, BLOCKS.Or, BLOCKS.While, BLOCKS.Array, BLOCKS.Get, BLOCKS.Set, BLOCKS.For, BLOCKS.Map, BLOCKS.Reduce, BLOCKS.Filter, BLOCKS.Push, BLOCKS.Pop, BLOCKS.Effect:
		return true
	default:
		return false
	}
}

type Specials struct {
	Plus        string
	Minus       string
	Multiply    string
	Divide      string
	Bang        string
	Equal       string
	LessThan    string
	GreaterThan string
	Ampersand   string
	Pipe        string
	Quote       string
	Dot         string
	Comment     string
	Eof         string
}

var SPECIALS = Specials{
	Plus:        "+",
	Minus:       "-",
	Multiply:    "*",
	Divide:      "/",
	Equal:       "=",
	Bang:        "!",
	LessThan:    "<",
	GreaterThan: ">",
	Ampersand:   "&",
	Pipe:        "|",
	Eof:         "EOF",
	Comment:     "#",
	Quote:       "\"",
	Dot:         ".",
}

func IsSpecial(text string) bool {
	switch text {
	case SPECIALS.Plus, SPECIALS.Minus, SPECIALS.Multiply, SPECIALS.Divide, SPECIALS.Bang, SPECIALS.LessThan, SPECIALS.GreaterThan, SPECIALS.Ampersand, SPECIALS.Pipe, SPECIALS.Eof:
		return true
	default:
		return false
	}
}

type Operators struct {
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
}

var OPERATORS = Operators{
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
}

func IsOperator(text string) bool {
	switch text {
	case OPERATORS.Plus, OPERATORS.Minus, OPERATORS.Multiply, OPERATORS.Divide, OPERATORS.Bang, OPERATORS.BangEqual, OPERATORS.EqualEqual, OPERATORS.LessThan, OPERATORS.LessThanEqual, OPERATORS.GreaterThan, OPERATORS.GreaterThanEqual, OPERATORS.Conjunction, OPERATORS.Disjunction:
		return true
	default:
		return false
	}
}

type Literals struct {
	Number     string
	String     string
	True       string
	False      string
	Nil        string
	Identifier string
}

var LITERALS = Literals{
	Number: "number",
	String: "string",
	True:   "true",
	False:  "false",
	Nil:    "nil",
}

func IsLiteral(text string) bool {
	switch text {
	case LITERALS.Number, LITERALS.String, LITERALS.True, LITERALS.False, LITERALS.Nil:
		return true
	default:
		return false
	}
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
	While            string
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
	Call:             "@",
	Val:              "val",
	Edit:             "let",
	If:               "if",
	Do:               "do",
	And:              "and",
	Or:               "or",
	While:            "while",
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

func IsToken(text string) bool {
	switch text {
	case TOKENS.Def, TOKENS.Call, TOKENS.Val, TOKENS.Edit, TOKENS.If, TOKENS.Do, TOKENS.And, TOKENS.Or, TOKENS.While, TOKENS.Array, TOKENS.Get, TOKENS.Set, TOKENS.For, TOKENS.Map, TOKENS.Reduce, TOKENS.Filter, TOKENS.Push, TOKENS.Pop, TOKENS.Effect, TOKENS.Plus, TOKENS.Minus, TOKENS.Multiply, TOKENS.Divide, TOKENS.Bang, TOKENS.BangEqual, TOKENS.EqualEqual, TOKENS.LessThan, TOKENS.LessThanEqual, TOKENS.GreaterThan, TOKENS.GreaterThanEqual, TOKENS.Conjunction, TOKENS.Disjunction, TOKENS.Number, TOKENS.String, TOKENS.Identifier, TOKENS.True, TOKENS.False, TOKENS.Nil, TOKENS.EndOfFile:
		return true
	default:
		return false
	}
}

type Identifiers struct {
	UpperA     string
	UpperB     string
	UpperC     string
	UpperD     string
	UpperE     string
	UpperF     string
	UpperG     string
	UpperH     string
	UpperI     string
	UpperJ     string
	UpperK     string
	UpperL     string
	UpperM     string
	UpperN     string
	UpperO     string
	UpperP     string
	UpperQ     string
	UpperR     string
	UpperS     string
	UpperT     string
	UpperU     string
	UpperV     string
	UpperW     string
	UpperX     string
	UpperY     string
	UpperZ     string
	LowerA     string
	LowerB     string
	LowerC     string
	LowerD     string
	LowerE     string
	LowerF     string
	LowerG     string
	LowerH     string
	LowerI     string
	LowerJ     string
	LowerK     string
	LowerL     string
	LowerM     string
	LowerN     string
	LowerO     string
	LowerP     string
	LowerQ     string
	LowerR     string
	LowerS     string
	LowerT     string
	LowerU     string
	LowerV     string
	LowerW     string
	LowerX     string
	LowerY     string
	LowerZ     string
	Underscore string
}

var IDENTIFIERS = Identifiers{
	UpperA:     "A",
	UpperB:     "B",
	UpperC:     "C",
	UpperD:     "D",
	UpperE:     "E",
	UpperF:     "F",
	UpperG:     "G",
	UpperH:     "H",
	UpperI:     "I",
	UpperJ:     "J",
	UpperK:     "K",
	UpperL:     "L",
	UpperM:     "M",
	UpperN:     "N",
	UpperO:     "O",
	UpperP:     "P",
	UpperQ:     "Q",
	UpperR:     "R",
	UpperS:     "S",
	UpperT:     "T",
	UpperU:     "U",
	UpperV:     "V",
	UpperW:     "W",
	UpperX:     "X",
	UpperY:     "Y",
	UpperZ:     "Z",
	LowerA:     "A",
	LowerB:     "B",
	LowerC:     "C",
	LowerD:     "D",
	LowerE:     "E",
	LowerF:     "F",
	LowerG:     "G",
	LowerH:     "H",
	LowerI:     "I",
	LowerJ:     "J",
	LowerK:     "K",
	LowerL:     "L",
	LowerM:     "M",
	LowerN:     "N",
	LowerO:     "O",
	LowerP:     "P",
	LowerQ:     "Q",
	LowerR:     "R",
	LowerS:     "S",
	LowerT:     "T",
	LowerU:     "U",
	LowerV:     "V",
	LowerW:     "W",
	LowerX:     "X",
	LowerY:     "Y",
	LowerZ:     "Z",
	Underscore: "_",
}

func IsIdentifier(text string) bool {
	switch text {
	case IDENTIFIERS.Underscore, IDENTIFIERS.LowerA, IDENTIFIERS.LowerB, IDENTIFIERS.LowerC, IDENTIFIERS.LowerD, IDENTIFIERS.LowerE, IDENTIFIERS.LowerF, IDENTIFIERS.LowerG, IDENTIFIERS.LowerH, IDENTIFIERS.LowerI, IDENTIFIERS.LowerJ, IDENTIFIERS.LowerK, IDENTIFIERS.LowerL, IDENTIFIERS.LowerM, IDENTIFIERS.LowerN, IDENTIFIERS.LowerO, IDENTIFIERS.LowerP, IDENTIFIERS.LowerQ, IDENTIFIERS.LowerR, IDENTIFIERS.LowerS, IDENTIFIERS.LowerT, IDENTIFIERS.LowerU, IDENTIFIERS.LowerV, IDENTIFIERS.LowerW, IDENTIFIERS.LowerX, IDENTIFIERS.LowerY, IDENTIFIERS.LowerZ, IDENTIFIERS.UpperA, IDENTIFIERS.UpperB, IDENTIFIERS.UpperC, IDENTIFIERS.UpperD, IDENTIFIERS.UpperE, IDENTIFIERS.UpperF, IDENTIFIERS.UpperG, IDENTIFIERS.UpperH, IDENTIFIERS.UpperI, IDENTIFIERS.UpperJ, IDENTIFIERS.UpperK, IDENTIFIERS.UpperL, IDENTIFIERS.UpperM, IDENTIFIERS.UpperN, IDENTIFIERS.UpperO, IDENTIFIERS.UpperP, IDENTIFIERS.UpperQ, IDENTIFIERS.UpperR, IDENTIFIERS.UpperS, IDENTIFIERS.UpperT, IDENTIFIERS.UpperU, IDENTIFIERS.UpperV, IDENTIFIERS.UpperW, IDENTIFIERS.UpperX, IDENTIFIERS.UpperY, IDENTIFIERS.UpperZ:
		return true
	default:
		return false
	}
}

type Numbers struct {
	Zero  string
	One   string
	Two   string
	Three string
	Four  string
	Five  string
	Six   string
	Seven string
	Eight string
	Nine  string
}

var NUMBERS = Numbers{
	Zero:  "0",
	One:   "1",
	Two:   "2",
	Three: "3",
	Four:  "4",
	Five:  "5",
	Six:   "6",
	Seven: "7",
	Eight: "8",
	Nine:  "9",
}

func IsNumber(text string) bool {
	switch text {
	case NUMBERS.Zero, NUMBERS.One, NUMBERS.Two, NUMBERS.Three, NUMBERS.Four, NUMBERS.Five, NUMBERS.Six, NUMBERS.Seven, NUMBERS.Eight, NUMBERS.Nine:
		return true
	default:
		return false
	}
}
