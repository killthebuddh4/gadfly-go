package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Colon(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
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
		return nil, errors.New("Schema is not a function")
	}

	parsed, err := schema(arg)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("Schema did not return a boolean")
	}

	return parsed, nil
}
