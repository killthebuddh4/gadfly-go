package eval

import (
	"errors"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func While(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	var value value.Value = nil

	for {
		condV, err := eval(trajectory.Children[0])

		if err != nil {
			return nil, err
		}

		cond, ok := condV.(bool)

		if !ok {
			return nil, errors.New("not a boolean")
		}

		if !cond {
			break
		}

		for _, child := range trajectory.Children[1:] {
			val, err := eval(child)

			if err != nil {
				return nil, err
			}

			value = val
		}
	}

	return value, nil
}
