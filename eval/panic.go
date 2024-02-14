package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Panic(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

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
