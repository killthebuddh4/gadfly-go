package eval

import (
	"errors"
	"fmt"

	"github.com/killthebuddh4/gadflai/types"
)

var Extract types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	base, ok := arguments[0].(map[string]types.Value)

	if !ok {
		return nil, errors.New("not a map")
	}

	keys, ok := arguments[1].([]types.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	extracted := make(map[string]types.Value)

	for _, keyV := range keys {
		key, ok := keyV.(string)

		if !ok {
			return nil, errors.New("key is not a string")
		}

		val, ok := base[key]

		if !ok {
			return nil, fmt.Errorf("key %s not found", key)
		}

		extracted[key] = val
	}

	return extracted, nil
}
