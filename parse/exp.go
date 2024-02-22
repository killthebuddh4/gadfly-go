package parse

import (
	"github.com/killthebuddh4/gadflai/types"
)

func GetExpDef(operator string) (types.ExpDef, error) {
	switch operator {
	case "if":
		return IF, nil
	case "std.write":
		return STD_WRITE, nil
	default:
		return EMPTY, nil
	}
}

var IF = types.ExpDef{
	Parameters: []types.Parameter{
		{Name: "if", IsThunk: false, EndWords: []string{"then"}},
		{Name: "then", IsThunk: true, EndWords: []string{"else"}},
		{Name: "else", IsThunk: true, EndWords: []string{"end", "catch"}},
	},
}

var EMPTY = types.ExpDef{
	Parameters: []types.Parameter{},
}

var STD_WRITE = types.ExpDef{
	Parameters: []types.Parameter{
		{Name: "std.write", IsThunk: false, EndWords: []string{"end", "catch"}},
	},
}
