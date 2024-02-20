package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var While types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	var value types.Value = nil

	for {
		condT, ok := arguments[0].(types.Thunk)

		if !ok {
			return nil, errors.New(":: While :: condition is not a thunk")
		}

		condV, err := condT()

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
			for _, arg := range arguments[1:] {
				bodyT, ok := arg.(types.Thunk)

				if !ok {
					return nil, errors.New(":: While :: body is not a thunk")
				}

				val, err := bodyT()

				if err != nil {
					return nil, err
				}

				value = val
			}
		}
	}

	return value, nil
}
