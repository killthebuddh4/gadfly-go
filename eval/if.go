package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var If types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	cond, ok := arguments[0].(bool)

	if !ok {
		return nil, errors.New(":: If :: condition is not a boolean")
	}

	if cond {
		body, ok := arguments[1].(types.Closure)

		if !ok {
			return nil, errors.New(":: If :: then is not a closure")
		}

		return body(scope)
	}

	if len(arguments) == 3 {
		body, ok := arguments[2].(types.Closure)

		if !ok {
			return nil, errors.New(":: If :: else is not a closure")
		}

		return body(scope)
	}

	return nil, nil
}
