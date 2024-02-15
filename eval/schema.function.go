package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func SchemaFunction() types.Exec {
	var lambda types.Exec = func(arguments ...types.Value) (types.Value, error) {
		raw := arguments[0]

		f, ok := raw.(types.Exec)

		if !ok {
			return nil, errors.New("SchemaFunction: not a function")
		}

		return f, nil
	}
	return lambda
}
