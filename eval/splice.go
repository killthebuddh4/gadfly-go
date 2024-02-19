package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Splice types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New("Splice :: not an array")
	}

	indexF, ok := arguments[1].(float64)

	if !ok {
		return nil, errors.New("Splice :: not a number")
	}

	index := int(indexF)

	if index < 0 {
		return nil, errors.New("Splice :: index cannot be negative")
	}

	if index > len(arr) {
		return nil, errors.New("Splice :: index cannot be greater than array length")
	}

	values, ok := arguments[2].([]types.Value)

	if !ok {
		return nil, errors.New("Splice :: not an array")
	}

	head := arr[:index]
	tail := arr[index:]

	spliced := append([]types.Value{}, head...)
	spliced = append(spliced, values...)
	spliced = append(spliced, tail...)

	return spliced, nil
}
