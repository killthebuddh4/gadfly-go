package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Reverse types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	types.ExpandTraj(scope)

	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New("Reverse :: not an array")
	}

	reversed := []types.Value{}

	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}

	return reversed, nil
}
