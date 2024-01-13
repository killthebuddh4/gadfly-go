package main

import (
	"errors"
	"reflect"
)

func EvaluateBangEqual(trajectory *Trajectory, args ...Value) (Value, error) {
	left, right := args[0], args[1]

	return left != right, nil
}

func EvaluateEqualEqual(trajectory *Trajectory, args ...Value) (Value, error) {
	left, right := args[0], args[1]

	return left == right, nil
}

func EvaluateGreaterThan(trajectory *Trajectory, args ...Value) (Value, error) {
	left, right := args[0], args[1]

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("left operand is not a number")
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return leftV > rightV, nil
}

func EvaluateGreaterThanEqual(trajectory *Trajectory, args ...Value) (Value, error) {
	left, right := args[0], args[1]

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("left operand is not a number")
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return leftV >= rightV, nil
}

func EvaluateLessThan(trajectory *Trajectory, args ...Value) (Value, error) {
	left, right := args[0], args[1]

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("left operand is not a number")
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return leftV < rightV, nil
}

func EvaluateLessThanEqual(trajectory *Trajectory, args ...Value) (Value, error) {
	left, right := args[0], args[1]

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("left operand is not a number")
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return leftV <= rightV, nil
}

func EvaluateMinus(trajectory *Trajectory, args ...Value) (Value, error) {
	if len(args) == 1 {
		return EvaluateMinusUnary(trajectory, args...)
	} else {
		return EvaluateMinusBinary(trajectory, args...)
	}
}

func EvaluateMinusBinary(trajectory *Trajectory, args ...Value) (Value, error) {
	left, right := args[0], args[1]

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("left operand is not a number")
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return leftV - rightV, nil
}

func EvaluateMinusUnary(trajectory *Trajectory, args ...Value) (Value, error) {
	right := args[0]

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return -rightV, nil
}

func EvaluatePlus(trajectory *Trajectory, args ...Value) (Value, error) {
	left, right := args[0], args[1]

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("left operand is not a number")
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return leftV + rightV, nil
}

func EvaluateDivide(trajectory *Trajectory, args ...Value) (Value, error) {
	left, right := args[0], args[1]

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("left operand is not a number")
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return leftV / rightV, nil
}

func EvaluateMultiply(trajectory *Trajectory, args ...Value) (Value, error) {
	left, right := args[0], args[1]

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("left operand is not a number " + reflect.TypeOf(left).String())
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number " + reflect.TypeOf(right).String())
	}

	return leftV * rightV, nil
}

func EvaluateBang(trajectory *Trajectory, args ...Value) (Value, error) {
	right := args[0]

	rightV, ok := right.(bool)

	if !ok {
		return nil, errors.New("right operand is not a boolean")
	}

	return !rightV, nil
}
