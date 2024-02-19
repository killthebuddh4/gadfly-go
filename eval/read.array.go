package eval

import (
	"errors"
	"fmt"

	"github.com/killthebuddh4/gadflai/types"
)

var ReadArray types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	if len(arguments) != 2 {
		return nil, errors.New("ReadArray :: wrong number of arguments")
	}

	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New(":: ReadArray :: not an array")
	}

	idx, ok := arguments[1].(float64)

	if !ok {
		return nil, errors.New(":: ReadArray :: not an integer")
	}

	if int(idx) < 0 || int(idx) >= len(arr) {
		return nil, fmt.Errorf(":: ReadArray :: index out of range: %d, length is %d", int(idx), len(arr))
	}

	return arr[int(idx)], nil
}
