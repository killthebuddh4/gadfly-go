package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Throw types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	signal, ok := arguments[0].(string)

	if !ok {
		return nil, errors.New("Throw :: signal not a string")
	}

	errHandlerV, err := types.ResolveError(scope.Parent, signal)

	if err != nil {
		return nil, err
	}

	errHandler, ok := errHandlerV.(types.Closure)

	if !ok {
		return nil, errors.New("Throw :: errHandler is not a function")
	}

	_, err = errHandler(scope, signal)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
