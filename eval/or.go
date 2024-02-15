package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Or types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	for _, arg := range arguments {
		child, ok := arg.(CaseHandler)

		if !ok {
			return nil, errors.New("not a case handler")
		}

		condV, err := child.Cond(scope)

		if err != nil {
			return nil, err
		}

		cond, ok := condV.(bool)

		if !ok {
			return nil, errors.New("not a boolean")
		}

		if cond {
			val, err := child.Body(scope)

			if err != nil {
				return nil, err
			}

			return val, nil
		}
	}
}
