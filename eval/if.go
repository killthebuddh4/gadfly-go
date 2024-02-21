package eval

import (
	"errors"
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

var If types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_EVAL")

	if debug {
		fmt.Println(":: If :: called")
	}

	if len(arguments) != 3 {
		return nil, fmt.Errorf(":: If :: requires 3 arguments but got %d", len(arguments))
	}

	condT, ok := arguments[0].(types.Thunk)

	if !ok {
		return nil, errors.New(":: If :: condition is not a thunk")
	}

	condV, err := condT()

	if err != nil {
		return nil, err
	}

	cond, ok := condV.(bool)

	if !ok {
		return nil, errors.New(":: If :: condition did not return a boolean")
	}

	if cond {
		body, ok := arguments[1].(types.Thunk)

		if !ok {
			return nil, errors.New(":: If :: then is not a thunk")
		}

		return body()
	} else {
		body, ok := arguments[2].(types.Thunk)

		if !ok {
			return nil, errors.New(":: If :: else is not a thunk")
		}

		return body()
	}
}
