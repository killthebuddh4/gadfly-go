package eval

import (
	"errors"
	"reflect"

	"github.com/killthebuddh4/gadflai/types"
)

var Emit types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	signal, ok := arguments[0].(string)

	if !ok {
		return nil, errors.New(":: Emit :: signal is not a string")
	}

	sigHandlerV, err := types.ResolveSignal(scope.Parent, signal)

	if err != nil {
		return nil, err
	}

	sigHandler, ok := sigHandlerV.(types.Closure)

	if !ok {
		return nil, errors.New(":: Emit :: sigHandler is not a function")
	}

	feedbackV, err := sigHandler(scope, signal)

	if err != nil {
		return nil, err
	}

	feedback, ok := feedbackV.(string)

	if !ok {
		return nil, errors.New(":: Emit :: feedback is not a string")
	}

	fbHandlers := []SignalHandler{}

	for _, fbHandlerV := range arguments[1:] {
		fbHandler, ok := fbHandlerV.(SignalHandler)

		if !ok {
			return nil, errors.New(":: Emit :: fbHandler is not a function , it's a " + reflect.TypeOf(fbHandlerV).String())
		}

		fbHandlers = append(fbHandlers, fbHandler)
	}

	for _, fbHandler := range fbHandlers {
		handler, err := fbHandler(feedback)

		if err != nil {
			continue
		}

		if !ok {
			return nil, errors.New(":: Emit :: fbHandler not a function")
		}

		return handler(scope, feedback)
	}

	return nil, errors.New(":: Emit :: no feedback handler found for signal")
}
