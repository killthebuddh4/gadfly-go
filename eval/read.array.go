package eval

import (
	"errors"
	"fmt"

	"github.com/killthebuddh4/gadflai/types"
)

func ReadArray(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	dataV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	data, ok := dataV.([]types.Value)

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
