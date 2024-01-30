package ai

import (
	"errors"
	"fmt"

	"github.com/killthebuddh4/gadflai/types"
)

func Call(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	fnVal, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	fn, ok := fnVal.(types.Lambda)

	if !ok {
		return nil, errors.New("Error evaluating call, expression that didn't evaluate to a Lambda, got " + fmt.Sprint(fnVal))
	}

	children := trajectory.Children[1:]

	args := []types.Value{}

	if len(children) > 0 {
		for _, traj := range children {
			arg, err := eval(traj)

			if err != nil {
				return nil, err
			}

			args = append(args, arg)
		}
	}

	val, err := fn(args...)

	if err != nil {
		return nil, err
	}

	return val, nil
}
