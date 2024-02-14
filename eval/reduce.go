package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Reduce(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	arrV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]types.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	initV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	fnV, err := eval(trajectory.Children[2])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(types.Lambda)

	if !ok {
		return nil, errors.New("not a function")
	}

	if (len(arr)) == 0 {
		return nil, nil
	}

	reduction := initV

	for i, v := range arr {
		reduction, err = fn(reduction, v, float64(i))

		if err != nil {
			return nil, err
		}
	}

	return reduction, nil
}
