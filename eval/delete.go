package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Delete(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	baseV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	base, ok := baseV.(map[string]types.Value)

	if !ok {
		return nil, errors.New("not a map")
	}

	keysV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	keys, ok := keysV.([]types.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	remainder := make(map[string]types.Value)

	for k, v := range base {
		var found bool = false
		for _, keyV := range keys {
			key, ok := keyV.(string)

			if !ok {
				return nil, errors.New("key is not a string")
			}

			if k == key {
				found = true
				break
			}
		}

		if !found {
			remainder[k] = v
		}
	}

	return remainder, nil
}
