package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func SchemaFunction() types.Lambda {
	var lambda types.Lambda = func(arguments ...types.Value) (types.Value, error) {
		raw := arguments[0]

		f, ok := raw.(types.Lambda)

		if !ok {
			return nil, errors.New("SchemaFunction: not a function")
		}

		return f, nil
	}
	return lambda
}
