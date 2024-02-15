package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Concat types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	result := ""

	for _, arg := range arguments {
		str, strOk := arg.(string)

		if !strOk {
			return nil, errors.New("Concat :: not a string")
		}

		result += str
	}

	return result, nil
}
