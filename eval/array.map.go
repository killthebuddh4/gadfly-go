package eval

import (
	"errors"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Map(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
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
		mapped, err := fn(v, float64(i))

		if err != nil {
			return nil, err
		}

		vals = append(vals, mapped)
	}

	return vals, nil
}
