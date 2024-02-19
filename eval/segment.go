package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Segment types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New("Segment :: first argument is not an array")
	}

	start, ok := arguments[1].(float64)

	if !ok {
		return nil, errors.New("Segment :: second argument is not an integer")
	}

	if start < 0 {
		return nil, errors.New("Segment :: start index cannot be negative")
	}

	end, ok := arguments[2].(float64)

	if !ok {
		return nil, errors.New("Segment :: third argument is not an integer")
	}

	if end < start {
		return nil, errors.New("Segment :: end index cannot be less than start index")
	}

	return append([]types.Value{}, arr[int(start):int(end)]...), nil
}
