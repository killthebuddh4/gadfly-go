package strings

import (
	"errors"

	"github.com/killthebuddh4/gadflai/eval"
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Split(trajectory *traj.Trajectory, eval eval.Eval) (value.Value, error) {
	traj.Expand(trajectory)

	arg, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	str, strOk := arg.(string)

	if !strOk {
		return nil, errors.New("split only accepts strings")
	}

	result := []value.Value{}

	for _, c := range str {
		result = append(result, string(c))
	}

	return result, nil
}
