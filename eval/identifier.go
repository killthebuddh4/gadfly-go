package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Identifier types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	if scope.Parent == nil {
		return nil, errors.New("cannot evaluate identifier " + scope.Expression.Operator.Value + " with nil parent")
	}
	return types.ResolveName(scope.Parent, scope.Expression.Operator.Value)
}
