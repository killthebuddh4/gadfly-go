package eval

import (
	"errors"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Splice(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	arrV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]value.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	indexV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	indexF, ok := indexV.(float64)

	if !ok {
		return nil, errors.New("not a number")
	}

	index := int(indexF)

	if index < 0 {
		return nil, errors.New("index cannot be negative")
	}

	if index > len(arr) {
		return nil, errors.New("index cannot be greater than array length")
	}

	valuesV, err := eval(trajectory.Children[2])

	if err != nil {
		return nil, err
	}

	values, ok := valuesV.([]value.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	head := arr[:index]
	tail := arr[index:]

	spliced := append([]value.Value{}, head...)
	spliced = append(spliced, values...)
	spliced = append(spliced, tail...)

	return spliced, nil
}
