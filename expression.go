package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Expression struct {
	Operator Token
	Inputs   []Expression
}

func PrintExpression(source string, exp Expression) {
	fmt.Println("----------------------------------------------")
	fmt.Println("OPERATOR")

	PrintToken(exp.Operator)

	fmt.Println("INPUTS")

	for _, input := range exp.Inputs {
		PrintToken(input.Operator)
	}

	fmt.Println("----------------------------------------------")

	for _, input := range exp.Inputs {
		PrintExpression(source, input)
	}
}

func Evaluate(exp Expression) (Value, error) {
	switch exp.Operator.Type {
	case "BANG_EQUAL":
		return EvaluateBangEqual(exp)
	case "EQUAL_EQUAL":
		return EvaluateEqualEqual(exp)
	case "GREATER":
		return EvaluateGreater(exp)
	case "GREATER_EQUAL":
		return EvaluateGreaterEqual(exp)
	case "LESS":
		return EvaluateLess(exp)
	case "LESS_EQUAL":
		return EvaluateLessEqual(exp)
	case "MINUS":
		return EvaluateMinus(exp)
	case "PLUS":
		return EvaluatePlus(exp)
	case "SLASH":
		return EvaluateSlash(exp)
	case "STAR":
		return EvaluateStar(exp)
	case "BANG":
		return EvaluateBang(exp)
	case "TRUE":
		return EvaluateTrue(exp)
	case "FALSE":
		return EvaluateFalse(exp)
	case "NIL":
		return EvaluateNil(exp)
	case "NUMBER":
		return EvaluateNumber(exp)
	case "STRING":
		return EvaluateString(exp)
	case "LEFT_PAREN":
		return EvaluateLeftParen(exp)
	default:
		return nil, errors.New("unknown operator")
	}
}

func EvaluateBangEqual(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

	if (leftErr != nil) || (rightErr != nil) {
		return nil, errors.New("error evaluating inputs")
	}

	return left != right, nil
}

func EvaluateEqualEqual(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

	if (leftErr != nil) || (rightErr != nil) {
		return nil, errors.New("error evaluating inputs")
	}

	return left == right, nil
}

func EvaluateGreater(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

	if (leftErr != nil) || (rightErr != nil) {
		return nil, errors.New("error evaluating inputs")
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

func EvaluateGreaterEqual(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

	if (leftErr != nil) || (rightErr != nil) {
		return nil, errors.New("error evaluating inputs")
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

func EvaluateLess(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

	if (leftErr != nil) || (rightErr != nil) {
		return nil, errors.New("error evaluating inputs")
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

func EvaluateLessEqual(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

	if (leftErr != nil) || (rightErr != nil) {
		return nil, errors.New("error evaluating inputs")
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

func EvaluateMinus(exp Expression) (Value, error) {
	if len(exp.Inputs) == 1 {
		return EvaluateMinusUnary(exp)
	} else {
		return EvaluateMinusBinary(exp)
	}
}

func EvaluateMinusBinary(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

	if (leftErr != nil) || (rightErr != nil) {
		return nil, errors.New("error evaluating inputs")
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

func EvaluateMinusUnary(exp Expression) (Value, error) {
	right, rightErr := Evaluate(exp.Inputs[1])

	if rightErr != nil {
		return nil, errors.New("error evaluating inputs")
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return -rightV, nil
}

func EvaluatePlus(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

	if (leftErr != nil) || (rightErr != nil) {
		return nil, errors.New("error evaluating inputs")
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

func EvaluateSlash(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

	if (leftErr != nil) || (rightErr != nil) {
		return nil, errors.New("error evaluating inputs")
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

func EvaluateStar(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

	if (leftErr != nil) || (rightErr != nil) {
		return nil, errors.New("error evaluating inputs")
	}

	leftV, ok := left.(float64)

	if !ok {
		return nil, errors.New("left operand is not a number")
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return leftV * rightV, nil
}

func EvaluateBang(exp Expression) (Value, error) {
	right, rightErr := Evaluate(exp.Inputs[1])

	if rightErr != nil {
		return nil, errors.New("error evaluating inputs")
	}

	rightV, ok := right.(bool)

	if !ok {
		return nil, errors.New("right operand is not a boolean")
	}

	return !rightV, nil
}

func EvaluateTrue(exp Expression) (Value, error) {
	return true, nil
}

func EvaluateFalse(exp Expression) (Value, error) {
	return false, nil
}

func EvaluateNil(exp Expression) (Value, error) {
	return nil, nil
}

func EvaluateNumber(exp Expression) (Value, error) {
	num, parseErr := strconv.ParseFloat(GetLexemeForToken(exp.Operator), 64)

	if parseErr != nil {
		return nil, errors.New("error parsing number")
	}

	return num, nil
}

func EvaluateString(exp Expression) (Value, error) {
	return GetLexemeForToken(exp.Operator), nil
}

func EvaluateLeftParen(exp Expression) (Value, error) {
	return nil, nil
}
