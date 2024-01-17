package array

import (
	"github.com/killthebuddh4/gadflai/eval"
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Array(trajectory *traj.Trajectory, eval eval.Eval) (value.Value, error) {
	traj.Expand(trajectory)

	arr := []value.Value{}

	for _, input := range trajectory.Children {
		val, err := eval(input)

		if err != nil {
			return nil, err
		}

		arr = append(arr, val)
	}

	return arr, nil
}
