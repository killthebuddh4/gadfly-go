package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var ReadHash types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	base, ok := arguments[0].(map[string]types.Value)

	if !ok {
		return nil, errors.New("ReadHash :: first argument is not a map")
	}

	key, ok := arguments[1].(string)

	if !ok {
		return nil, errors.New("ReadHash :: second argument is not a string")
	}

	val, ok := base[key]

	if !ok {
		return nil, nil
	} else {
		return val, nil
	}
}
