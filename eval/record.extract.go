package eval

import (
	"errors"
	"fmt"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Extract(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	baseV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	base, ok := baseV.(map[string]value.Value)

	if !ok {
		return nil, errors.New("not a record")
	}

	keysV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	keys, ok := keysV.([]value.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	extracted := make(map[string]value.Value)

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
