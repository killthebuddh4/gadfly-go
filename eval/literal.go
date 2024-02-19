package eval

import (
	"errors"
	"strconv"
	"strings"

	"github.com/killthebuddh4/gadflai/types"
)

var True types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	return true, nil
}

var False types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	return false, nil
}

var Nil types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	return nil, nil
}

var Number types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	num, parseErr := strconv.ParseFloat(scope.Expression.Operator.Value, 64)

	if parseErr != nil {
		return nil, errors.New("error parsing number")
	}

	return num, nil
}

var String types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	return strings.Trim(scope.Expression.Operator.Value, "\""), nil
}
