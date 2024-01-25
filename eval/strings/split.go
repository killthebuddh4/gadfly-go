package strings

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Split(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	arg, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	str, strOk := arg.(string)

	if !strOk {
		return nil, errors.New("split only accepts strings")
	}

	result := []types.Value{}

	for _, c := range str {
		result = append(result, string(c))
	}

	return result, nil
}
