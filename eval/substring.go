package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Substring types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	str, ok := arguments[0].(string)

	if !ok {
		return nil, errors.New(":: Substring :: not a string")
	}

	start, ok := arguments[1].(float64)

	if !ok {
		return nil, errors.New(":: Substring :: not a number")
	}

	end, ok := arguments[2].(float64)

	if !ok {
		return nil, errors.New(":: Substring :: not a number")
	}

	return str[int(start):int(end)], nil
}
