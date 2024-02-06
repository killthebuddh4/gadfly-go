package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Pop(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	arrV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]types.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	return arr[:len(arr)-1], nil
}
