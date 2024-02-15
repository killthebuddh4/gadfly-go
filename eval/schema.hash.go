package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

func SchemaHash() types.Exec {
	var lambda types.Exec = func(arguments ...types.Value) (types.Value, error) {
		raw := arguments[0]

		hash, ok := raw.(map[string]types.Value)

		if !ok {
			return nil, errors.New("not a hash")
		}

		return hash, nil
	}

	return lambda
}
