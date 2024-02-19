package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Find types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New(":: Find :: not an array")
	}

	fn, ok := arguments[1].(types.Exec)

	if !ok {
		return nil, errors.New(":: Find :: not a function")
	}

	for i, v := range arr {
		foundV, err := fn(scope, v, float64(i))

		if err != nil {
			return nil, err
		}

		found, ok := foundV.(bool)

		if !ok {
			return nil, errors.New(":: Find :: found is not a boolean")
		}

		if found {
			return v, nil
		}
	}

	return nil, nil
}
