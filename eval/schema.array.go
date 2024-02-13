package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func SchemaArray() types.Lambda {
	var lambda types.Lambda = func(arguments ...types.Value) (types.Value, error) {
		raw := arguments[0]

		array, ok := raw.([]types.Value)

		if !ok {
			return nil, errors.New("not an array")
		}

		return array, nil
	}

	return lambda
}
