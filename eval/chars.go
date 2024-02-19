package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Chars types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	str, ok := arguments[0].(string)

	if !ok {
		return nil, errors.New("Chars :: chars only accepts strings")
	}

	result := []types.Value{}

	for _, c := range str {
		result = append(result, string(c))
	}

	return result, nil
}
