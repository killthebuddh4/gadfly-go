package exec

import (
	"github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

type Evaluator func(*trajectory.Trajectory) (value.Value, error)

func Exec(trajectory *trajectory.Trajectory) (value.Value, error) {
	eval, dispatchErr := dispatch(trajectory)

	if dispatchErr != nil {
		return nil, dispatchErr
	}

	val, evalErr := eval(trajectory, Exec)

	if evalErr != nil {
		return nil, evalErr
	}

	return val, nil
}
