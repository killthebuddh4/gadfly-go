package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var And types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	var value types.Value = nil

	for _, arg := range arguments {
		child, ok := arg.(types.Thunk)

		if !ok {
			return nil, errors.New(":: And :: argument is not a thunk")
		}

		val, err := child()

		if err != nil {
			return nil, err
		}

		if val == nil {
			return nil, nil
		} else {
			value = val
		}
	}

	return value, nil
}
