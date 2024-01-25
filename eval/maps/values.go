package maps

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Values(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	baseV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	base, ok := baseV.(map[string]types.Value)

	if !ok {
		return nil, errors.New("not a map")
	}

	values := []types.Value{}

	for _, v := range base {
		values = append(values, v)
	}

	return values, nil
}
