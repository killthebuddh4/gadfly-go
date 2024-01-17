package eval

import (
	"errors"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Find(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
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

	for i, v := range arr {
		foundV, err := fn(v, float64(i))

		if err != nil {
			return nil, err
		}

		found, ok := foundV.(bool)

		if !ok {
			return nil, errors.New("found is not a boolean")
		}

		if found {
			return v, nil
		}
	}

	return nil, nil
}
