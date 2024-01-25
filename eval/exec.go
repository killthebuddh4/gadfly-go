package eval

import (
	"github.com/killthebuddh4/gadflai/types"
)

type Evaluator func(*types.Trajectory) (types.Value, error)

func Eval(trajectory *types.Trajectory) (types.Value, error) {
	eval, dispatchErr := dispatch(trajectory)

	if dispatchErr != nil {
		return nil, dispatchErr
	}

	val, evalErr := eval(trajectory, Eval)

	if evalErr != nil {
		return nil, evalErr
	}

	return val, nil
}
