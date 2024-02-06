package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Hash(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	if (len(trajectory.Children) % 2) != 0 {
		return nil, errors.New("map must have even number of inputs")
	}

	maps := make(map[string]types.Value)

	for i := 0; i < len(trajectory.Children); i += 2 {
		keyVal, err := eval(trajectory.Children[i])

		if err != nil {
			return nil, err
		}

		key, ok := keyVal.(string)

		if !ok {
			return nil, errors.New("key is not a string")
		}

		valV, err := eval(trajectory.Children[i+1])

		if err != nil {
			return nil, err
		}

		maps[key] = valV
	}

	return maps, nil
}
