package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Signal types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	// HACK this is a hack this is a hack this is a hack
	identifier := scope.Children[0].Expression.Operator.Value

	handler, ok := arguments[1].(types.Closure)

	if !ok {
		return nil, errors.New("Signal :: handler not a function")
	}

	types.DefineSignal(scope.Parent, identifier, handler)

	return handler, nil
}
