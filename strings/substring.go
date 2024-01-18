package strings

import (
	"errors"

	"github.com/killthebuddh4/gadflai/eval"
	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Substring(trajectory *traj.Trajectory, eval eval.Eval) (value.Value, error) {
	traj.Expand(trajectory)

	arg, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	str, strOk := arg.(string)

	if !strOk {
		return nil, errors.New("split only accepts strings")
	}

	startV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	start, startOk := startV.(float64)

	if !startOk {
		return nil, errors.New("start index must be a number")
	}

	endV, err := eval(trajectory.Children[2])

	if err != nil {
		return nil, err
	}

	end, endOk := endV.(float64)

	if !endOk {
		return nil, errors.New("end index must be a number")
	}

	return str[int(start):int(end)], nil
}
