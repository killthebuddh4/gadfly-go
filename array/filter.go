package array

import (
	"errors"

	"github.com/killthebuddh4/gadflai/eval"
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Filter(trajectory *traj.Trajectory, eval eval.Eval) (value.Value, error) {
	traj.Expand(trajectory)

	arrV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]value.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	fnV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(traj.Lambda)

	if !ok {
		return nil, errors.New("not a function")
	}

	vals := []value.Value{}

	for i, v := range arr {
		filterV, err := fn(v, float64(i))

		if err != nil {
			return nil, err
		}

		filter, ok := filterV.(bool)

		if !ok {
			return nil, errors.New("filter is not a boolean")
		}

		if filter {
			vals = append(vals, v)
		}
	}

	return vals, nil
}
