package eval

import (
	"github.com/killthebuddh4/gadflai/types"
)

func Do(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	var val types.Value

	for _, input := range trajectory.Children {
		v, err := eval(input)

		val = v

		if err != nil {
			return nil, err
		}
	}

	return val, nil
}
