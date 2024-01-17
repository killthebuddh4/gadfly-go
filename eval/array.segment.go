package eval

import (
	"errors"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Segment(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	arrV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]value.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	startV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	start, ok := startV.(float64)

	if !ok {
		return nil, errors.New("not an integer")
	}

	if start < 0 {
		return nil, errors.New("start index cannot be negative")
	}

	endV, err := eval(trajectory.Children[2])

	if err != nil {
		return nil, err
	}

	end, ok := endV.(float64)

	if !ok {
		return nil, errors.New("not an integer")
	}

	if end < start {
		return nil, errors.New("end index cannot be less than start index")
	}

	return append([]value.Value{}, arr[int(start):int(end)]...), nil
}
