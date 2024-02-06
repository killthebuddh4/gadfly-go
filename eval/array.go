package eval

import (
	"github.com/killthebuddh4/gadflai/types"
)

func Array(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	arr := []types.Value{}

	for _, input := range trajectory.Children {
		val, err := eval(input)

		if err != nil {
			return nil, err
		}

		arr = append(arr, val)
	}

	return arr, nil
}
