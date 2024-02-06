package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func ReadHash(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	baseV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	base, ok := baseV.(map[string]types.Value)

	if !ok {
		return nil, errors.New("not a map")
	}

	keyV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	key, ok := keyV.(string)

	if !ok {
		return nil, errors.New("not a string")
	}

	val, ok := base[key]

	if !ok {
		return nil, nil
	} else {
		return val, nil
	}
}
