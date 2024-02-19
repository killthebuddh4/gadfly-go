package eval

import (
	"errors"
	"fmt"

	"github.com/killthebuddh4/gadflai/types"
)

var WriteArray types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New("WriteArray :: error setting array, data is not an array, it is " + fmt.Sprint(arr))
	}

	index, ok := arguments[1].(float64)

	if !ok {
		return nil, errors.New("WriteArray :: error setting array, index is not a number, it is " + fmt.Sprint(index))
	}

	result := make([]types.Value, len(arr))

	for i, v := range arr {
		if float64(i) == index {
			result[i] = arguments[2]
		} else {
			result[i] = v
		}
	}

	return result, nil
}
