package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Push types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New("Push :: first argument is not an array")
	}

	val := arguments[1]

	arr = append(arr, val)

	return arr, nil
}
