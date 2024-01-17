package eval

import (
	"errors"
	"fmt"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Get(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	dataV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	data, ok := dataV.([]value.Value)

	if !ok {
		return nil, errors.New("error getting from array, data is not an array, it is " + fmt.Sprint(dataV))
	}

	indexV, err := eval(trajectory.Children[1])

	index, ok := indexV.(float64)

	if !ok {
		return nil, errors.New("error getting from array, index is not a number, it is " + fmt.Sprint(indexV))
	}

	if err != nil {
		return nil, err
	}

	val := data[int(index)]

	return val, nil
}
