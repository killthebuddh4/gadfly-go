package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Values types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	base, ok := arguments[0].(map[string]types.Value)

	if !ok {
		return nil, errors.New("not a map")
	}

	values := []types.Value{}

	for _, v := range base {
		values = append(values, v)
	}

	return values, nil
}
