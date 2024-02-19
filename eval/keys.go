package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Keys types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	base, ok := arguments[0].(map[string]types.Value)

	if !ok {
		return nil, errors.New("not a map")
	}

	keys := []types.Value{}

	for k := range base {
		keys = append(keys, k)
	}

	return keys, nil
}
