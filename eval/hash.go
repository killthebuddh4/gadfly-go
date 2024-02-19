package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Hash types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	if (len(arguments) % 2) != 0 {
		return nil, errors.New("Hash :: map must have even number of inputs")
	}

	hash := make(map[string]types.Value)

	for i := 0; i < len(arguments); i += 2 {
		key, ok := arguments[i].(string)

		if !ok {
			return nil, errors.New("Hash :: key is not a string")
		}

		hash[key] = arguments[i+1]
	}

	return hash, nil
}
