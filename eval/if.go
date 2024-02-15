package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var If types.Exec = func(arguments ...types.Value) (types.Value, error) {
	cond, ok := arguments[0].(bool)

	if !ok {
		return nil, errors.New("condition is not a boolean")
	}

	if cond {
		body, ok := arguments[1].(types.Exec)

		if !ok {
			return nil, errors.New("body is not a lambda")
		}

		return body()
	}

	if len(arguments) == 3 {
		body, ok := arguments[2].(types.Exec)

		if !ok {
			return nil, errors.New("body is not a lambda")
		}

		return body()
	}

	return nil, nil
}
