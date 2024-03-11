package eval

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/killthebuddh4/gadflai/types"
)

var When types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	cond, ok := arguments[0].(bool)

	if !ok {
		return nil, errors.New(":: When :: condition is not a boolean")
	}

	fmt.Println("arguments[1],", reflect.TypeOf(arguments[1]))

	body, ok := arguments[1].(types.Thunk)

	if !ok {
		return nil, errors.New(":: When :: body is not a Thunk")
	}

	if cond {
		return body()
	}

	return nil, nil
}
