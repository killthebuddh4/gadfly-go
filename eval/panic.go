package eval

import (
	"errors"

	traj "github.com/killthebuddh4/gadflai/trajectory"
	"github.com/killthebuddh4/gadflai/value"
)

func Panic(trajectory *traj.Trajectory, eval Eval) (value.Value, error) {
	traj.Expand(trajectory)

	messageV, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	message, ok := messageV.(string)

	if !ok {
		return nil, errors.New("panic message must be a string")
	}

	panic(message)
}
