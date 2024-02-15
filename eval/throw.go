package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Throw(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	errorV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	errorSym, ok := errorV.(string)

	if !ok {
		return nil, errors.New("not a string")
	}

	errHandlerV, err := types.ResolveError(trajectory.Parent, errorSym)

	if err != nil {
		return nil, err
	}

	errHandler, ok := errHandlerV.(types.Exec)

	if !ok {
		return nil, errors.New("Throw :: errHandler is not a function")
	}

	_, err = errHandler(errorSym)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
