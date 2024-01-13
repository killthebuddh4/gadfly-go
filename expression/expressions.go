package expression

type Expressions struct {
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

var EXPRESSIONS = Expressions{
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

func IsExpressionType(text string) bool {
	switch text {
	case EXPRESSIONS.Fn, EXPRESSIONS.Pipe, EXPRESSIONS.Def, EXPRESSIONS.Call, EXPRESSIONS.Val, EXPRESSIONS.Edit, EXPRESSIONS.If, EXPRESSIONS.Do, EXPRESSIONS.And, EXPRESSIONS.Or, EXPRESSIONS.While, EXPRESSIONS.Array, EXPRESSIONS.Get, EXPRESSIONS.Set, EXPRESSIONS.For, EXPRESSIONS.Map, EXPRESSIONS.Reduce, EXPRESSIONS.Filter, EXPRESSIONS.Push, EXPRESSIONS.Pop, EXPRESSIONS.End, EXPRESSIONS.Effect, EXPRESSIONS.Plus, EXPRESSIONS.Minus, EXPRESSIONS.Multiply, EXPRESSIONS.Divide, EXPRESSIONS.Bang, EXPRESSIONS.BangEqual, EXPRESSIONS.EqualEqual, EXPRESSIONS.LessThan, EXPRESSIONS.LessThanEqual, EXPRESSIONS.GreaterThan, EXPRESSIONS.GreaterThanEqual, EXPRESSIONS.Number, EXPRESSIONS.String, EXPRESSIONS.Identifier, EXPRESSIONS.True, EXPRESSIONS.False, EXPRESSIONS.Nil, EXPRESSIONS.EndOfFile:
		return true
	default:
		return false
	}
}
