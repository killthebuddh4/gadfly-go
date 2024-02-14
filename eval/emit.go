package eval

import (
	"errors"
	"reflect"

	"github.com/killthebuddh4/gadflai/types"
)

func Emit(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	signalV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	signal, ok := signalV.(string)

	if !ok {
		return nil, errors.New("not a string")
	}

	sigHandlerV, err := types.ResolveSignal(trajectory.Parent, signal)

	if err != nil {
		return nil, err
	}

	sigHandler, ok := sigHandlerV.(types.Lambda)

	if !ok {
		return nil, errors.New("Emit :: sigHandler is not a function")
	}

	feedbackV, err := sigHandler(signal)

	if err != nil {
		return nil, err
	}

	feedback, ok := feedbackV.(string)

	if !ok {
		return nil, errors.New("Emit :: feedback is not a string")
	}

	fbHandlers := []SignalHandler{}

	for _, child := range trajectory.Children[1:] {
		fbHandlerV, err := eval(child)

		if err != nil {
			return nil, err
		}

		fbHandler, ok := fbHandlerV.(SignalHandler)

		if !ok {
			return nil, errors.New("Emit :: fbHandler is not a function , it's a " + reflect.TypeOf(fbHandlerV).String())
		}

		fbHandlers = append(fbHandlers, fbHandler)
	}

	for _, fbHandler := range fbHandlers {
		handler, err := fbHandler(feedback)

		if err != nil {
			continue
		}

		if !ok {
			return nil, errors.New("not a function")
		}

		return handler(feedbackV)
	}

	return nil, errors.New("no feedback handler found for signal")
}
