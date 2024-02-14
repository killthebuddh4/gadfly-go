package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func Schema(trajectory *types.Trajectory, eval types.Eval) (types.Value, error) {
	types.ExpandTraj(trajectory)

	schemas := []types.Lambda{}

	for _, child := range trajectory.Children {
		schemaV, err := eval(child)

		if err != nil {
			return nil, err
		}

		schema, ok := schemaV.(types.Lambda)

		if !ok {
			return nil, errors.New("not a function")
		}

		schemas = append(schemas, schema)
	}

	var lambda types.Lambda = func(arguments ...types.Value) (types.Value, error) {
		raw := arguments[0]

		var val types.Value = raw

		for _, schema := range schemas {
			v, err := schema(val)

			if err != nil {
				return nil, err
			}

			val = v
		}

		return val, nil
	}

	return lambda, nil
}
