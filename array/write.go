package array

import (
	"errors"
	"fmt"

	"github.com/killthebuddh4/gadflai/eval"
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Write(trajectory *traj.Trajectory, eval eval.Eval) (value.Value, error) {
	traj.Expand(trajectory)

	dataV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	data, ok := dataV.([]value.Value)

	if !ok {
		return nil, errors.New("error setting array, data is not an array, it is " + fmt.Sprint(dataV))
	}

	indexV, err := eval(trajectory.Children[1])

	index, ok := indexV.(float64)

	if !ok {
		return nil, errors.New("error setting array, index is not a number, it is " + fmt.Sprint(indexV))
	}

	if err != nil {
		return nil, err
	}

	val, err := eval(trajectory.Children[2])

	if err != nil {
		return nil, err
	}

	result := make([]value.Value, len(data))

	for i, v := range data {
		if float64(i) == index {
			result[i] = val
		} else {
			result[i] = v
		}
	}

	return result, nil
}
