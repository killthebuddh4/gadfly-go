package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Shift(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	arrV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]types.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	return append([]types.Value{}, arr[1:]...), nil
}
