package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var And types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	var value types.Value = nil

	for _, arg := range arguments {
		child, ok := arg.(CaseHandler)

		if !ok {
			return nil, errors.New("And :: not a case handler")
		}

		condV, err := child.Cond(scope, nil)

		if err != nil {
			return nil, err
		}

		cond, ok := condV.(bool)

		if !ok {
			return nil, errors.New("And :: not a boolean")
		}

		if !cond {
			return nil, nil
		} else {
			val, err := child.Body(scope, nil)

			if err != nil {
				return nil, err
			}

			value = val
		}
	}

	return value, nil
}
