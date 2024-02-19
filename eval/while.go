package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var While types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	var value types.Value = nil

	for {
		cond, ok := arguments[0].(types.Closure)

		if !ok {
			return nil, errors.New(":: While :: condition is not a boolean")
		}

		condV, err := cond(scope)

		if err != nil {
			return nil, err
		}

		cont, ok := condV.(bool)

		if !ok {
			return nil, errors.New(":: While :: condition is not a boolean")
		}

		if !cont {
			break
		} else {
			body, ok := arguments[1].(types.Closure)

			if !ok {
				return nil, errors.New(":: While :: body is not a lambda")
			}

			val, err := body(scope)

			if err != nil {
				return nil, err
			}

			value = val
		}
	}

	return value, nil
}
