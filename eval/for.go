package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func For(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
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
		_, err := fn(v, float64(i))

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
