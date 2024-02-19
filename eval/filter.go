package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Filter types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New("Filter :: not an array")
	}

	fn, ok := arguments[1].(types.Closure)

	if !ok {
		return nil, errors.New("Filter :: not a function")
	}

	vals := []types.Value{}

	for i, v := range arr {
		filterV, err := fn(scope, v, float64(i))

		if err != nil {
			return nil, err
		}

		filter, ok := filterV.(bool)

		if !ok {
			return nil, errors.New("Filter :: filter result is not a boolean")
		}

		if filter {
			vals = append(vals, v)
		}
	}

	return vals, nil
}
