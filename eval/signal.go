package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Signal(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	identifier := trajectory.Children[0].Expression.Operator.Value

	var handler types.Exec

	handlerV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	handler, ok := handlerV.(types.Exec)

	if !ok {
		return nil, errors.New("not a function")
	}

	types.DefineSignal(trajectory.Parent, identifier, handler)

	return handler, nil
}
