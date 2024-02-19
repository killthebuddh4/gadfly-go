package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Merge types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	base, ok := arguments[0].(map[string]types.Value)

	if !ok {
		return nil, errors.New("not a map")
	}

	new, ok := arguments[1].(map[string]types.Value)

	if !ok {
		return nil, errors.New("not a map")
	}

	merged := make(map[string]types.Value)

	for k, v := range base {
		merged[k] = v
	}

	for k, v := range new {
		merged[k] = v
	}

	return merged, nil
}
