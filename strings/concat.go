package strings

import (
	"errors"

	"github.com/killthebuddh4/gadflai/eval"
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Concat(trajectory *traj.Trajectory, eval eval.Eval) (value.Value, error) {
	traj.Expand(trajectory)

	result := ""

	for _, child := range trajectory.Children {
		arg, err := eval(child)

		if err != nil {
			return nil, err
		}

		str, strOk := arg.(string)

		if !strOk {
			return nil, errors.New("concat only accepts strings")
		}

		result += str
	}

	return result, nil
}
