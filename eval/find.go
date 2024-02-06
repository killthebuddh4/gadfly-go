package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Find(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	arrV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]types.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	fnV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(types.Lambda)

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
