package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var Def types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	name, ok := arguments[0].(string)

	if !ok {
		return nil, errors.New(":: Def :: identifier is not a string")
	}

	if len(arguments) < 2 {
		return nil, errors.New(":: Def :: no value to define")
	}

	value := arguments[len(arguments)-1]

	err := types.DefineName(scope.Parent, name, value)

	if err != nil {
		return nil, err
	}

	return value, nil
}
