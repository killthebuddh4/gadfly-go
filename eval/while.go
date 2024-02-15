package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var While types.Exec = func(arguments ...types.Value) (types.Value, error) {
	var value types.Value = nil

	for {
		cond, ok := arguments[0].(bool)

		if !ok {
			return nil, errors.New("condition is not a boolean")
		}

		if !cond {
			break
		} else {
			body, ok := arguments[1].(types.Exec)

			if !ok {
				return nil, errors.New("body is not a lambda")
			}

			val, err := body()

			if err != nil {
				return nil, err
			}

			value = val
		}
	}

	return value, nil
}
