package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Chars(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	arg, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	str, strOk := arg.(string)

	if !strOk {
		return nil, errors.New("chars only accepts strings")
	}

	result := []types.Value{}

	for _, c := range str {
		result = append(result, string(c))
	}

	return result, nil
}
