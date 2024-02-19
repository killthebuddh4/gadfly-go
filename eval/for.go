package eval

import (
	"errors"
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
)

var For types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	_, debug := os.LookupEnv("GADFLY_DEBUG_EVAL")

	if debug {
		fmt.Println(":: For :: called")
	}

	arr, ok := arguments[0].([]types.Value)

	if !ok {
		return nil, errors.New(":: For :: not an array")
	}

	fn, ok := arguments[1].(types.Closure)

	if !ok {
		return nil, errors.New(":: For :: not a closure")
	}

	for i, v := range arr {
		_, err := fn(scope, v, float64(i))

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
