package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

type CaseHandler struct {
	Cond types.Closure
	Body types.Closure
}

var Case types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	cond, ok := arguments[0].(types.Closure)

	if !ok {
		return nil, errors.New("condition is not a lambda")
	}

	body, ok := arguments[1].(types.Closure)

	if !ok {
		return nil, errors.New("body is not a lambda")
	}

	handler := CaseHandler{Cond: cond, Body: body}

	return handler, nil
}
