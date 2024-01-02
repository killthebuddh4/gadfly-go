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

type Evaluator func(Expression) (Value, error)

func Evaluate(exp Expression) (Value, error) {
	var eval Evaluator
	switch exp.Operator.Type {
	case "BANG_EQUAL":
		eval = EvaluateBangEqual
	case "EQUAL_EQUAL":
		eval = EvaluateEqualEqual
	case "GREATER":
		eval = EvaluateGreater
	case "GREATER_EQUAL":
		eval = EvaluateGreaterEqual
	case "LESS":
		eval = EvaluateLess
	case "LESS_EQUAL":
		eval = EvaluateLessEqual
	case "MINUS":
		eval = EvaluateMinus
	case "PLUS":
		eval = EvaluatePlus
	case "SLASH":
		eval = EvaluateSlash
	case "STAR":
		eval = EvaluateStar
	case "BANG":
		eval = EvaluateBang
	case "TRUE":
		eval = EvaluateTrue
	case "FALSE":
		eval = EvaluateFalse
	case "NIL":
		eval = EvaluateNil
	case "NUMBER":
		eval = EvaluateNumber
	case "STRING":
		eval = EvaluateString
	case "LEFT_PAREN":
		eval = EvaluateLeftParen
	case "IDENTIFIER":
		eval = EvaluateIdentifier
	case "let":
		eval = EvaluateLet
	case "do", "then", "else":
		eval = EvaluateDo
	case "if":
		eval = EvaluateIf
	default:
		return nil, errors.New("unknown operator " + exp.Operator.Type)
	}

	return eval(exp)
}

func EvaluateBangEqual(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	return left != right, nil
}

func EvaluateEqualEqual(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	return left == right, nil
}

func EvaluateGreater(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

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

func EvaluateGreaterEqual(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

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

func EvaluateLess(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

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

func EvaluateLessEqual(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

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

func EvaluateMinusUnary(exp Expression) (Value, error) {
	right, rightErr := Evaluate(exp.Inputs[1])

	if rightErr != nil {
		return nil, rightErr
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

func EvaluateSlash(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

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

func EvaluateLet(exp Expression) (Value, error) {
	identifier := GetLexemeForToken(exp.Inputs[0].Operator)

	val, err := Evaluate(exp.Inputs[1])

	if err != nil {
		return nil, err
	}

	setSymbolErr := SetSymbol(identifier, val)

	if setSymbolErr != nil {
		return nil, setSymbolErr
	}

	return val, nil
}

func EvaluateIdentifier(exp Expression) (Value, error) {
	identifier := GetLexemeForToken(exp.Operator)
	return GetSymbol(identifier)
}

func EvaluateLeftParen(exp Expression) (Value, error) {
	return nil, nil
}

func EvaluateDo(exp Expression) (Value, error) {
	PushEnvironment()

	var val Value

	for _, input := range exp.Inputs {
		v, err := Evaluate(input)

		val = v

		if err != nil {
			PopEnvironment()
			return nil, err
		}
	}

	PopEnvironment()

	return val, nil
}

func EvaluateIf(exp Expression) (Value, error) {
	condition, err := Evaluate(exp.Inputs[0])

	if err != nil {
		return nil, err
	}

	conditionVal, ok := condition.(bool)

	if !ok {
		return nil, errors.New("condition is not a boolean")
	}

	if conditionVal {
		return Evaluate(exp.Inputs[1])
	} else {
		return Evaluate(exp.Inputs[2])
	}
}
