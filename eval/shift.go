package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Shift types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New("Shift :: not an array")
	}

	return append([]types.Value{}, arr[1:]...), nil
}
