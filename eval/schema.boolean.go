package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func SchemaBoolean() types.Lambda {
	var lambda types.Lambda = func(arguments ...types.Value) (types.Value, error) {
		raw := arguments[0]

		val, ok := raw.(bool)

		if !ok {
			return nil, errors.New("not a boolean")
		}

		return val, nil
	}

	return lambda
}
