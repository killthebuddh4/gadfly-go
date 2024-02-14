package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func WriteHash(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
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

	valV, err := eval(trajectory.Children[2])

	if err != nil {
		return nil, err
	}

	written := make(map[string]types.Value)

	for k, v := range base {
		written[k] = v
	}

	written[key] = valV

	return written, nil
}
