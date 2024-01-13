package main

import (
	"errors"
	"fmt"
	"reflect"
)

func EvaluateBangEqual(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	if len(args) != 2 {
		return nil, errors.New("bang equal must have two arguments")
	}

	left, right := args[0], args[1]

	return left != right, nil
}

func EvaluateEqualEqual(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	if len(args) != 2 {
		return nil, errors.New("equal equal must have two arguments")
	}

	left, right := args[0], args[1]

	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	fmt.Println("EVALUATING EQUAL EQUAL", left, right)
	return left == right, nil
}

func EvaluateGreaterThan(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	if len(args) != 2 {
		return nil, errors.New("greater than must have two arguments")
	}

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

func EvaluateGreaterThanEqual(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	if len(args) != 2 {
		return nil, errors.New("greater than equal must have two arguments")
	}

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

func EvaluateLessThan(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	if len(args) != 2 {
		return nil, errors.New("less than must have two arguments")
	}

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

func EvaluateLessThanEqual(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	if len(args) != 2 {
		return nil, errors.New("less than equal must have two arguments")
	}

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

func EvaluateMinus(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	if len(args) == 1 {
		return EvaluateMinusUnary(scope)
	} else {
		return EvaluateMinusBinary(scope)
	}
}

func EvaluateMinusBinary(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

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

func EvaluateMinusUnary(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	right := args[0]

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return -rightV, nil
}

func EvaluatePlus(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

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

func EvaluateDivide(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

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

func EvaluateMultiply(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

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

func EvaluateBang(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	right := args[0]

	rightV, ok := right.(bool)

	if !ok {
		return nil, errors.New("right operand is not a boolean")
	}

	return !rightV, nil
}

func EvaluateConjunction(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}

	for _, arg := range args {
		argV, ok := arg.(bool)

		if !ok {
			return nil, errors.New("not a boolean")
		}

		if !argV {
			return false, nil
		}
	}

	return true, nil
}

func EvaluateDisjunction(scope *Trajectory) (Value, error) {
	expand(scope)

	args := []Value{}

	for _, child := range scope.Children {
		arg, err := evaluate(child)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}

	for _, arg := range args {
		argV, ok := arg.(bool)

		if !ok {
			return nil, errors.New("not a boolean")
		}

		if argV {
			return true, nil
		}
	}

	return false, nil
}
