package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Catch types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	identifier, ok := arguments[0].(string)

	if !ok {
		return nil, errors.New("Catch :: identifier is not a string")
	}

	handler, ok := arguments[1].(types.Exec)

	if !ok {
		return nil, errors.New("Catch :: handler is not a lambda")
	}

	types.DefineError(scope.Parent, identifier, handler)

	return handler, nil
}
