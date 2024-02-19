package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Unshift types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	head := []types.Value{arguments[1]}

	return append(head, arr...), nil
}
