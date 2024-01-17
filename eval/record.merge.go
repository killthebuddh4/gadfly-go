package eval

import (
	"errors"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Merge(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	baseV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	base, ok := baseV.(map[string]value.Value)

	if !ok {
		return nil, errors.New("not a record")
	}

	newV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	new, ok := newV.(map[string]value.Value)

	if !ok {
		return nil, errors.New("not a record")
	}

	merged := make(map[string]value.Value)

	for k, v := range base {
		merged[k] = v
	}

	for k, v := range new {
		merged[k] = v
	}

	return merged, nil
}
