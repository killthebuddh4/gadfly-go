package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Unshift(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	arrV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]types.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	val, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	head := []types.Value{val}

	return append(head, arr...), nil
}
