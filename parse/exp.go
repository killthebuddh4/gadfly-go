package parse

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func GetEndwords(operator types.Operator) ([]string, error) {
	switch operator.Type {
	case "if":
		return []string{"then"}, nil
	case "then":
		return []string{"else"}, nil
	case "else":
		return []string{"catch", "end"}, nil
	case "do":
		return []string{"catch", "end"}, nil
	case "std.write":
		return []string{"catch", "end"}, nil
	default:
		return nil, errors.New(":: GetSpec :: Unknown operator type: " + operator.Type)
	}
}
