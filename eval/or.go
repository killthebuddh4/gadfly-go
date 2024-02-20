package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Or types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	for _, arg := range arguments {
		child, ok := arg.(types.Thunk)

		if !ok {
			return nil, errors.New(":: Or :: argument is not a thunk")
		}

		val, err := child()

		if err != nil {
			return nil, err
		}

		if val != nil {
			return val, nil
		}
	}

	return nil, nil
}
