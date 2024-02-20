package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var When types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	condT, ok := arguments[0].(types.Thunk)

	if !ok {
		return nil, errors.New(":: When :: condition is not a lambda")
	}

	body, ok := arguments[1].(types.Thunk)

	if !ok {
		return nil, errors.New(":: When :: body is not a lambda")
	}

	condV, err := condT()

	if err != nil {
		return nil, err
	}

	cond, ok := condV.(bool)

	if !ok {
		return nil, errors.New(":: When :: condition did not return a boolean")
	}

	if cond {
		return body()
	}

	return nil, nil
}
