package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Delete types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	base, ok := arguments[0].(map[string]types.Value)

	if !ok {
		return nil, errors.New("Delete :: not a map")
	}

	keys, ok := arguments[1].([]types.Value)

	if !ok {
		return nil, errors.New("Delete :: not an array")
	}

	remainder := make(map[string]types.Value)

	for k, v := range base {
		var found bool = false
		for _, keyV := range keys {
			key, ok := keyV.(string)

			if !ok {
				return nil, errors.New("Delete :: key is not a string")
			}

			if k == key {
				found = true
				break
			}
		}

		if !found {
			remainder[k] = v
		}
	}

	return remainder, nil
}
