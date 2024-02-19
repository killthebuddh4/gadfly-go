package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Reduce types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New("Reduce :: not an array")
	}

	init := arguments[1]

	fn, ok := arguments[2].(types.Closure)

	if !ok {
		return nil, errors.New("Reduce :: not a function")
	}

	if (len(arr)) == 0 {
		return init, nil
	}

	reduction := init

	for i, v := range arr {
		next, err := fn(scope, reduction, v, float64(i))

		if err != nil {
			return nil, err
		}

		reduction = next
	}

	return reduction, nil
}
