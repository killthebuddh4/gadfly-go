package record

import (
	"errors"

	"github.com/killthebuddh4/gadflai/eval"
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Keys(trajectory *traj.Trajectory, eval eval.Eval) (value.Value, error) {
	traj.Expand(trajectory)

	baseV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	base, ok := baseV.(map[string]value.Value)

	if !ok {
		return nil, errors.New("not a record")
	}

	keys := []value.Value{}

	for k := range base {
		keys = append(keys, k)
	}

	return keys, nil
}
