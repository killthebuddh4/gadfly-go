package eval

import (
	"errors"

	"github.com/killthebuddh4/gadflai/types"
)

var BangEqual types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	return arguments[0] != arguments[1], nil
}

var EqualEqual types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	return arguments[0] == arguments[1], nil
}

var Greater types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	left, ok := arguments[0].(float64)

	if !ok {
		return nil, errors.New("Greater :: left operand is not a number")
	}

	right, ok := arguments[1].(float64)

	if !ok {
		return nil, errors.New("Greater :: right operand is not a number")
	}

	return left > right, nil
}

var GreaterEqual types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	left, ok := arguments[0].(float64)

	if !ok {
		return nil, errors.New("GreaterEqual :: left operand is not a number")
	}

	right, ok := arguments[1].(float64)

	if !ok {
		return nil, errors.New("GreaterEqual :: right operand is not a number")
	}

	return left >= right, nil
}

var Less types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	left, ok := arguments[0].(float64)

	if !ok {
		return nil, errors.New("Less :: left operand is not a number")
	}

	right, ok := arguments[1].(float64)

	if !ok {
		return nil, errors.New("Less :: right operand is not a number")
	}

	return left < right, nil
}

var LessEqual types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	left, ok := arguments[0].(float64)

	if !ok {
		return nil, errors.New("LessEqual :: left operand is not a number")
	}

	right, ok := arguments[1].(float64)

	if !ok {
		return nil, errors.New("LessEqual :: right operand is not a number")
	}

	return left <= right, nil
}

var Minus types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	if len(arguments) == 1 {
		return -arguments[0].(float64), nil
	} else {
		return arguments[0].(float64) - arguments[1].(float64), nil
	}
}

var Plus types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	left, ok := arguments[0].(float64)

	if !ok {
		return nil, errors.New("Plus :: left operand is not a number")
	}

	right, ok := arguments[1].(float64)

	if !ok {
		return nil, errors.New("Plus :: right operand is not a number")
	}

	return left + right, nil
}

var Multiply types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	left, ok := arguments[0].(float64)

	if !ok {
		return nil, errors.New("Multiply :: left operand is not a number")
	}

	right, ok := arguments[1].(float64)

	if !ok {
		return nil, errors.New("Multiply :: right operand is not a number")
	}

	return left * right, nil
}

var Divide types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	left, ok := arguments[0].(float64)

	if !ok {
		return nil, errors.New("Divide :: left operand is not a number")
	}

	right, ok := arguments[1].(float64)

	if !ok {
		return nil, errors.New("Divide :: right operand is not a number")
	}

	return left / right, nil
}

var Bang types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	operand, ok := arguments[0].(bool)

	if !ok {
		return nil, errors.New("Bang :: operand is not a boolean")
	}

	return !operand, nil
}

var Conjunction types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	left, ok := arguments[0].(bool)

	if !ok {
		return nil, errors.New("Conjunction :: left operand is not a boolean")
	}

	if !left {
		return false, nil
	}

	right, ok := arguments[1].(bool)

	if !ok {
		return nil, errors.New("Conjunction :: right operand is not a boolean")
	}

	return right, nil
}

var Disjunction types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	left, ok := arguments[0].(bool)

	if !ok {
		return nil, errors.New("Disjunction :: left operand is not a boolean")
	}

	if left {
		return true, nil
	}

	right, ok := arguments[1].(bool)

	if !ok {
		return nil, errors.New("Disjunction :: right operand is not a boolean")
	}

	return right, nil
}
