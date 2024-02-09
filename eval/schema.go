package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Schema(trajectory *types.Trajectory, eval types.Exec) (types.Value, error) {
	types.ExpandTraj(trajectory)

	arg, err := eval(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	schemaV, err := eval(trajectory.Children[1])

	if err != nil {
		return nil, err
	}

	schema, ok := schemaV.(types.Lambda)

	if !ok {
		return nil, errors.New("not a function")
	}

	return schema(arg)
}
