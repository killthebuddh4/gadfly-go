package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func SchemaNumber() types.Lambda {
	var lambda types.Lambda = func(arguments ...types.Value) (types.Value, error) {
		raw := arguments[0]

		num, ok := raw.(float64)

		if !ok {
			return nil, errors.New("not a number")
		}

		return num, nil
	}

	return lambda
}
