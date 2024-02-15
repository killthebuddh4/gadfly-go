package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func SchemaString() types.Exec {
	var lambda types.Exec = func(arguments ...types.Value) (types.Value, error) {
		raw := arguments[0]

		str, ok := raw.(string)

		if !ok {
			return nil, errors.New("not a string")
		}

		return str, nil
	}

	return lambda
}
