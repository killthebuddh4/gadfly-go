package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

type SignalHandler func(string) (types.Closure, error)

var On types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	signal, ok := arguments[0].(string)

	if !ok {
		return nil, errors.New("On :: not a string")
	}

	handlerBody, ok := arguments[1].(types.Closure)

	if !ok {
		return nil, errors.New("On :: not a function")
	}

	var handler SignalHandler = func(dispatched string) (types.Closure, error) {
		if dispatched != signal {
			return nil, errors.New("signal mismatch")
		} else {
			return handlerBody, nil
		}
	}

	return handler, nil
}
