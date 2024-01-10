package main

import (
	"errors"
	"reflect"
	"strconv"
)

type Evaluator func(*Trajectory) (Value, error)

func Evaluate(trajectory *Trajectory) (Value, error) {
	eval, dispatchErr := dispatch(trajectory)

	if dispatchErr != nil {
		return nil, dispatchErr
	}

	val, evalErr := eval(trajectory)

	if evalErr != nil {
		return nil, evalErr
	}

	return val, nil
}

func EvaluateBangEqual(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	left, leftErr := Evaluate(trajectory.Children[0])
	right, rightErr := Evaluate(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	return left != right, nil
}

func EvaluateEqualEqual(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	left, leftErr := Evaluate(trajectory.Children[0])
	right, rightErr := Evaluate(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	return left == right, nil
}

func EvaluateGreater(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	left, leftErr := Evaluate(trajectory.Children[0])
	right, rightErr := Evaluate(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

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

func EvaluateGreaterEqual(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	left, leftErr := Evaluate(trajectory.Children[0])
	right, rightErr := Evaluate(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

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

func EvaluateLess(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	left, leftErr := Evaluate(trajectory.Children[0])
	right, rightErr := Evaluate(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

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

func EvaluateLessEqual(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	left, leftErr := Evaluate(trajectory.Children[0])
	right, rightErr := Evaluate(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

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

func EvaluateMinus(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	if len(trajectory.Children) == 1 {
		return EvaluateMinusUnary(trajectory)
	} else {
		return EvaluateMinusBinary(trajectory)
	}
}

func EvaluateMinusBinary(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	left, leftErr := Evaluate(trajectory.Children[0])
	right, rightErr := Evaluate(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

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

func EvaluateMinusUnary(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	right, rightErr := Evaluate(trajectory.Children[1])

	if rightErr != nil {
		return nil, rightErr
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return -rightV, nil
}

func EvaluatePlus(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	left, leftErr := Evaluate(trajectory.Children[0])
	right, rightErr := Evaluate(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

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

func EvaluateSlash(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	left, leftErr := Evaluate(trajectory.Children[0])
	right, rightErr := Evaluate(trajectory.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

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

func EvaluateStar(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	left, leftErr := Evaluate(trajectory.Children[0])
	right, rightErr := Evaluate(trajectory.Children[1])

	if (leftErr != nil) || (rightErr != nil) {
		return nil, errors.New("error evaluating inputs")
	}

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

func EvaluateBang(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	right, rightErr := Evaluate(trajectory.Children[1])

	if rightErr != nil {
		return nil, errors.New("error evaluating inputs")
	}

	rightV, ok := right.(bool)

	if !ok {
		return nil, errors.New("right operand is not a boolean")
	}

	return !rightV, nil
}

func EvaluateTrue(trajectory *Trajectory) (Value, error) {
	return true, nil
}

func EvaluateFalse(trajectory *Trajectory) (Value, error) {
	return false, nil
}

func EvaluateNil(trajectory *Trajectory) (Value, error) {
	return nil, nil
}

func EvaluateNumber(trajectory *Trajectory) (Value, error) {
	num, parseErr := strconv.ParseFloat(GetLexemeForToken(trajectory.Expression.Operator), 64)

	if parseErr != nil {
		return nil, errors.New("error parsing number")
	}

	return num, nil
}

func EvaluateString(trajectory *Trajectory) (Value, error) {
	return trajectory.Expression.Operator.Lexeme, nil
}

func EvaluateSet(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	dataV, err := Evaluate(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	data, ok := dataV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	indexV, err := Evaluate(trajectory.Children[1])

	index, ok := indexV.(float64)

	if !ok {
		return nil, errors.New("not a number")
	}

	if err != nil {
		return nil, err
	}

	val, err := Evaluate(trajectory.Children[2])

	if err != nil {
		return nil, err
	}

	data[int(index)] = val

	return data, nil
}

func EvaluateGet(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	dataV, err := Evaluate(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	data, ok := dataV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	indexV, err := Evaluate(trajectory.Children[1])

	index, ok := indexV.(float64)

	if !ok {
		return nil, errors.New("not a number")
	}

	if err != nil {
		return nil, err
	}

	val := data[int(index)]

	return val, nil
}

func EvaluateDo(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	var val Value

	for _, input := range trajectory.Children {
		v, err := Evaluate(input)

		val = v

		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

func EvaluateLogical(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	left, leftErr := Evaluate(trajectory.Children[0])

	if leftErr != nil {
		return nil, leftErr
	}

	leftV, ok := left.(bool)

	if !ok {
		return nil, errors.New("left operand is not a boolean")
	}

	if trajectory.Expression.Operator.Type == TOKENS.Conjunction {
		if !leftV {
			return false, nil
		}
	} else if trajectory.Expression.Operator.Type == TOKENS.Disjunction {
		if leftV {
			return true, nil
		}
	} else {
		return nil, errors.New("unknown logical operator, && and || are supported")
	}

	right, rightErr := Evaluate(trajectory.Children[1])

	if rightErr != nil {
		return nil, rightErr
	}

	rightV, ok := right.(bool)

	if !ok {
		return nil, errors.New("right operand is not a boolean")
	}

	return rightV, nil
}
