package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Segment(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	arrV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]types.Value)

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

	return append([]types.Value{}, arr[int(start):int(end)]...), nil
}
