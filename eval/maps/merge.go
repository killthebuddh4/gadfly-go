package maps

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Merge(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	baseV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	base, ok := baseV.(map[string]types.Value)

	if !ok {
		return nil, errors.New("not a map")
	}

	newV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	new, ok := newV.(map[string]types.Value)

	if !ok {
		return nil, errors.New("not a map")
	}

	merged := make(map[string]types.Value)

	for k, v := range base {
		merged[k] = v
	}

	for k, v := range new {
		merged[k] = v
	}

	return merged, nil
}
