package eval

import (
	"errors"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Push(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	arrV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]value.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	val, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	arr = append(arr, val)

	return arr, nil
}
