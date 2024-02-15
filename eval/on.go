package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

type SignalHandler func(string) (types.Exec, error)

func On(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	signalExpr := trajectory.Children[0]

	signalV, err := eval(signalExpr)

	if err != nil {
		return nil, err
	}

	signal, ok := signalV.(string)

	if !ok {
		return nil, errors.New("not a string")
	}

	handlerExpr := trajectory.Children[1]

	handlerV, err := eval(handlerExpr)

	if err != nil {
		return nil, err
	}

	handlerBody, ok := handlerV.(types.Exec)

	if !ok {
		return nil, errors.New("not a function")
	}

	var handler SignalHandler = func(dispatched string) (types.Exec, error) {
		if dispatched != signal {
			return nil, errors.New("signal mismatch")
		} else {
			return handlerBody, nil
		}
	}

	return handler, nil
}
