package eval

import (
	"errors"
	"fmt"

	"github.com/killthebuddh4/gadflai/types"
)

var Signal types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	if len(arguments) != 2 {
		return nil, fmt.Errorf(":: Signal :: expected 2 arguments, got %d", len(arguments))
	}

	// HACK this is a hack this is a hack this is a hack
	identifier := scope.Expression.Parameters[0].Operator.Value

	handler, ok := arguments[1].(types.Closure)

	if !ok {
		return nil, errors.New(":: Signal :: handler not a function")
	}

	types.DefineSignal(scope.Parent, identifier, handler)

	return handler, nil
}
