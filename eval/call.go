package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Call types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	fn, ok := arguments[0].(types.Closure)

	if !ok {
		return nil, errors.New(":: Call :: first argument is not a function")
	}

	return fn(scope, arguments[1:]...)
}
