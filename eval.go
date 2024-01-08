package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Evaluator func(Expression) (Value, error)

func Evaluate(exp Expression) (Value, error) {
	fmt.Println("Evaluating")
	var eval Evaluator
	switch exp.Operator.Type {
	case "ROOT":
		eval = func(exp Expression) (Value, error) {
			var value Value
			for _, input := range exp.Inputs {
				val, err := Evaluate(input)

				if err != nil {
					return nil, err
				}

				value = val
			}

			return value, nil
		}
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
		fmt.Println("Called STAR in evaluate")
		if exp.Parent != nil {
			fmt.Println("Exp.parent type ", exp.Parent.Operator.Type)
		}

		eval = EvaluateStar
	case "BANG":
		eval = EvaluateBang
	case "true":
		eval = EvaluateTrue
	case "false":
		eval = EvaluateFalse
	case "nil":
		eval = EvaluateNil
	case "NUMBER":
		eval = EvaluateNumber
	case "STRING":
		eval = EvaluateString
	case "IDENTIFIER":
		eval = EvaluateIdentifier
	case "def":
		eval = func(exp Expression) (Value, error) {
			return EvaluateDef(exp.Parent, exp)
		}
	case "call":
		eval = EvaluateCall
	case "val":
		eval = EvaluateVal
	case "let":
		eval = EvaluateLet
	case "filter":
		eval = EvaluateFilter
	case "for":
		eval = EvaluateFor
	case "map":
		eval = EvaluateMap
	case "reduce":
		eval = EvaluateReduce
	case "do", "when", "then", "else":
		eval = EvaluateDo
	case "and":
		eval = EvaluateAnd
	case "or":
		eval = EvaluateOr
	case "fn":
		eval = EvaluateLambda
	case "array":
		eval = EvaluateArray
	case "set":
		eval = EvaluateSet
	case "get":
		eval = EvaluateGet
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
	fmt.Println("Evaluating STAR of type", exp.Operator.Type)
	fmt.Println("Num inmputs", len(exp.Inputs))
	fmt.Println("Type", exp.Operator.Type)
	left, leftErr := Evaluate(exp.Inputs[0])
	right, rightErr := Evaluate(exp.Inputs[1])

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

func EvaluateVal(exp Expression) (Value, error) {
	identifier := GetLexemeForToken(exp.Inputs[0].Operator)

	val, err := Evaluate(exp.Inputs[1])

	if err != nil {
		return nil, err
	}

	setSymbolErr := DefSymbol(identifier, val)

	if setSymbolErr != nil {
		return nil, setSymbolErr
	}

	return val, nil
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

func EvaluateSet(exp Expression) (Value, error) {
	dataV, err := Evaluate(exp.Inputs[0])

	if err != nil {
		return nil, err
	}

	data, ok := dataV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	indexV, err := Evaluate(exp.Inputs[1])

	index, ok := indexV.(float64)

	if !ok {
		return nil, errors.New("not a number")
	}

	if err != nil {
		return nil, err
	}

	val, err := Evaluate(exp.Inputs[2])

	if err != nil {
		return nil, err
	}

	data[int(index)] = val

	return data, nil
}

func EvaluateGet(exp Expression) (Value, error) {
	dataV, err := Evaluate(exp.Inputs[0])

	if err != nil {
		return nil, err
	}

	data, ok := dataV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	indexV, err := Evaluate(exp.Inputs[1])

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

func EvaluateLambda(exp Expression) (Value, error) {
	var lambda Lambda = func(arguments ...Value) (Value, error) {
		PushEnvironment()

		parameters := exp.Inputs[0]

		if len(arguments) != len(parameters.Inputs) {
			return nil, errors.New("wrong number of arguments")
		}

		if len(parameters.Inputs) > 0 {
			for i, parameter := range parameters.Inputs {
				identifier := GetLexemeForToken(parameter.Operator)

				val := arguments[i]

				setSymbolErr := DefSymbol(identifier, val)

				if setSymbolErr != nil {
					return nil, setSymbolErr
				}
			}
		}

		body := exp.Inputs[1]

		val, err := Evaluate(body)

		if err != nil {
			return nil, err
		}

		PopEnvironment()

		return val, nil
	}

	return lambda, nil
}

func EvaluateDef(parent *Expression, exp Expression) (Value, error) {
	if parent == nil {
		return nil, errors.New("def must be inside a scope")
	}

	identifier := GetLexemeForToken(exp.Inputs[0].Operator)

	var value Value

	for _, input := range exp.Inputs[1:] {
		val, err := Evaluate(input)

		if err != nil {
			return nil, err
		}

		value = val
	}

	lambda, ok := value.(Lambda)

	if !ok {
		return nil, errors.New("def body must be a function")
	}

	setDefinition(parent, identifier, lambda)

	return lambda, nil
}

func EvaluateCall(exp Expression) (Value, error) {
	PushEnvironment()

	lambda, err := getDefinition(&exp, GetLexemeForToken(exp.Operator))

	if err != nil {
		return nil, err
	}

	argsExps := exp.Inputs

	args := []Value{}

	if len(argsExps) > 0 {
		for _, argExp := range argsExps {
			arg, err := Evaluate(argExp)

			if err != nil {
				return nil, err
			}

			args = append(args, arg)
		}
	}

	val, err := lambda(args...)

	if err != nil {
		return nil, err
	}

	PopEnvironment()

	fmt.Println("Call result", val)
	return val, nil
}

func EvaluateIdentifier(exp Expression) (Value, error) {
	identifier := GetLexemeForToken(exp.Operator)
	return GetSymbol(identifier)
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
	whenExp := exp.Inputs[0]
	thenExp := exp.Inputs[1]
	elseExp := exp.Inputs[2]

	conditionVal, err := Evaluate(whenExp)

	if err != nil {
		return nil, err
	}

	condition, ok := conditionVal.(bool)

	if !ok {
		return nil, errors.New("condition is not a boolean")
	}

	if condition {
		return Evaluate(thenExp)
	} else {
		return Evaluate(elseExp)
	}
}

func EvaluateFor(exp Expression) (Value, error) {
	arrV, err := Evaluate(exp.Inputs[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	fnV, err := Evaluate(exp.Inputs[1])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(func(args ...Value) (Value, error))

	if !ok {
		return nil, errors.New("not a function")
	}

	for i, v := range arr {
		_, err := fn(v, float64(i))

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func EvaluateMap(exp Expression) (Value, error) {
	arrV, err := Evaluate(exp.Inputs[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	fnV, err := Evaluate(exp.Inputs[1])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(func(args ...Value) (Value, error))

	if !ok {
		return nil, errors.New("not a function")
	}

	vals := []Value{}

	for i, v := range arr {
		fmt.Println("Evaluating map", v, float64(i))
		mapped, err := fn(v, float64(i))

		if err != nil {
			return nil, err
		}

		vals = append(vals, mapped)
	}

	return vals, nil
}

func EvaluateFilter(exp Expression) (Value, error) {
	arrV, err := Evaluate(exp.Inputs[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	fnV, err := Evaluate(exp.Inputs[1])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(func(args ...Value) (Value, error))

	if !ok {
		return nil, errors.New("not a function")
	}

	vals := []Value{}

	for i, v := range arr {
		filterV, err := fn(v, float64(i))

		if err != nil {
			return nil, err
		}

		filter, ok := filterV.(bool)

		if !ok {
			return nil, errors.New("filter is not a boolean")
		}

		if filter {
			vals = append(vals, v)
		}
	}

	return vals, nil
}

func EvaluateReduce(exp Expression) (Value, error) {
	arrV, err := Evaluate(exp.Inputs[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	initV, err := Evaluate(exp.Inputs[1])

	if err != nil {
		return nil, err
	}

	fnV, err := Evaluate(exp.Inputs[2])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(func(args ...Value) (Value, error))

	if !ok {
		return nil, errors.New("not a function")
	}

	if (len(arr)) == 0 {
		return nil, nil
	}

	reduction := initV

	for i, v := range arr {
		reduction, err = fn(reduction, v, float64(i))

		if err != nil {
			return nil, err
		}
	}

	return reduction, nil
}

func EvaluateArray(exp Expression) (Value, error) {
	arr := []Value{}

	for _, input := range exp.Inputs {
		val, err := Evaluate(input)

		if err != nil {
			return nil, err
		}

		arr = append(arr, val)
	}

	return arr, nil
}

func EvaluateAnd(exp Expression) (Value, error) {
	if (len(exp.Inputs) % 2) != 0 {
		return nil, errors.New("and must have even number of inputs")
	}

	var val Value = nil

	for i := 0; i < len(exp.Inputs); i += 2 {
		conditionVal, err := Evaluate(exp.Inputs[i])

		if err != nil {
			return nil, err
		}

		condition, ok := conditionVal.(bool)

		if !ok {
			return nil, errors.New("condition is not a boolean")
		}

		if !condition {
			return false, nil
		}

		body, err := Evaluate(exp.Inputs[i+1])

		if err != nil {
			return nil, err
		}

		val = body
	}

	return val, nil
}

func EvaluateOr(exp Expression) (Value, error) {
	if (len(exp.Inputs) % 2) != 0 {
		return nil, errors.New("or must have even number of inputs")
	}

	for i := 0; i < len(exp.Inputs); i += 2 {
		conditionVal, err := Evaluate(exp.Inputs[i])

		if err != nil {
			return nil, err
		}

		condition, ok := conditionVal.(bool)

		if !ok {
			return nil, errors.New("condition is not a boolean")
		}

		if condition {
			return Evaluate(exp.Inputs[i+1])
		}
	}

	return nil, nil
}

func EvaluateLogical(exp Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Inputs[0])

	if leftErr != nil {
		return nil, leftErr
	}

	leftV, ok := left.(bool)

	if !ok {
		return nil, errors.New("left operand is not a boolean")
	}

	if exp.Operator.Type == "and" {
		if !leftV {
			return false, nil
		}
	} else if exp.Operator.Type == "or" {
		if leftV {
			return true, nil
		}
	}

	right, rightErr := Evaluate(exp.Inputs[1])

	if rightErr != nil {
		return nil, rightErr
	}

	rightV, ok := right.(bool)

	if !ok {
		return nil, errors.New("right operand is not a boolean")
	}

	return rightV, nil
}
