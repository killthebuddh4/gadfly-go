package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Keys(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	baseV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	base, ok := baseV.(map[string]types.Value)

	if !ok {
		return nil, errors.New("not a map")
	}

	keys := []types.Value{}

	for k := range base {
		keys = append(keys, k)
	}

	return keys, nil
}
