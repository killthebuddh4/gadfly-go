package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Colon types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	schema, ok := arguments[1].(types.Closure)

	if !ok {
		return nil, errors.New("Colon :: Schema is not a function")
	}

	return schema(scope, arguments[0])
}
