package record

import (
	"errors"

	"github.com/killthebuddh4/gadflai/eval"
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Read(trajectory *traj.Trajectory, eval eval.Eval) (value.Value, error) {
	traj.Expand(trajectory)

	baseV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	base, ok := baseV.(map[string]value.Value)

	if !ok {
		return nil, errors.New("not a record")
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
