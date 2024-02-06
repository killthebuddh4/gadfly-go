package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Map(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
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

	vals := []types.Value{}

	for i, v := range arr {
		mapped, err := fn(v, float64(i))

		if err != nil {
			return nil, err
		}

		vals = append(vals, mapped)
	}

	return vals, nil
}
