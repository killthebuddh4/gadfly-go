package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Map types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New(":: Map :: not an array")
	}

	fn, ok := arguments[1].(types.Closure)

	if !ok {
		return nil, errors.New(":: Map :: not a function")
	}

	vals := []types.Value{}

	for i, v := range arr {
		mapped, err := fn(scope, v, float64(i))

		if err != nil {
			return nil, err
		}

		vals = append(vals, mapped)
	}

	return vals, nil
}
