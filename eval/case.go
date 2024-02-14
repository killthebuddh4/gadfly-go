package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

type CaseHandler struct {
	Cond types.Lambda
	Body types.Lambda
}

func Case(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	condV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	cond, ok := condV.(types.Lambda)

	if !ok {
		return nil, errors.New("not a function")
	}

	bodyV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	body, ok := bodyV.(types.Lambda)

	if !ok {
		return nil, errors.New("not a function")
	}

	handler := CaseHandler{Cond: cond, Body: body}

	return handler, nil
}
