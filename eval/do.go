package eval

import (
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Do(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	var val value.Value

	for _, input := range trajectory.Children {
		v, err := eval(input)

		val = v

		if err != nil {
			return nil, err
		}
	}

	return val, nil
}
