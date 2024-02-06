package eval

import (
	"errors"
	"fmt"

	"github.com/killthebuddh4/gadflai/types"
)

func Extract(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
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

	extracted := make(map[string]types.Value)

	for _, keyV := range keys {
		key, ok := keyV.(string)

		if !ok {
			return nil, errors.New("key is not a string")
		}

		val, ok := base[key]

		if !ok {
			return nil, fmt.Errorf("key %s not found", key)
		}

		extracted[key] = val
	}

	return extracted, nil
}
