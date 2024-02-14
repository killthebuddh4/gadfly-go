package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Concat(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

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
