package main

import (
	"errors"
	"reflect"
	"strconv"
)

type Evaluator func(*Expression) (Value, error)

func Evaluate(exp *Expression) (Value, error) {
	var eval Evaluator
	switch exp.Operator.Type {
	case "ROOT":
		eval = func(exp *Expression) (Value, error) {
			var value Value
			for _, input := range exp.Children {
				val, err := Evaluate(input)

				if err != nil {
					return nil, err
				}

				value = val
			}

			return value, nil
		}
	case TOKENS.BangEqual:
		eval = EvaluateBangEqual
	case TOKENS.EqualEqual:
		eval = EvaluateEqualEqual
	case TOKENS.GreaterThan:
		eval = EvaluateGreater
	case TOKENS.GreaterThanEqual:
		eval = EvaluateGreaterEqual
	case TOKENS.LessThan:
		eval = EvaluateLess
	case TOKENS.LessThanEqual:
		eval = EvaluateLessEqual
	case TOKENS.Conjunction, TOKENS.Disjunction:
		eval = EvaluateLogical
	case TOKENS.Minus:
		eval = EvaluateMinus
	case TOKENS.Plus:
		eval = EvaluatePlus
	case TOKENS.Divide:
		eval = EvaluateSlash
	case TOKENS.Multiply:
		eval = EvaluateStar
	case TOKENS.Bang:
		eval = EvaluateBang
	case TOKENS.True:
		eval = EvaluateTrue
	case TOKENS.False:
		eval = EvaluateFalse
	case TOKENS.Nil:
		eval = EvaluateNil
	case TOKENS.Number:
		eval = EvaluateNumber
	case TOKENS.String:
		eval = EvaluateString
	case TOKENS.Identifier:
		eval = EvaluateIdentifier
	case TOKENS.Def:
		eval = EvaluateDef
	case TOKENS.Call:
		eval = EvaluateCall
	case TOKENS.Edit:
		eval = EvaluateLet
	case TOKENS.Filter:
		eval = EvaluateFilter
	case TOKENS.For:
		eval = EvaluateFor
	case TOKENS.Map:
		eval = EvaluateMap
	case TOKENS.Reduce:
		eval = EvaluateReduce
	case TOKENS.Push:
		eval = EvaluatePush
	case TOKENS.Pop:
		eval = EvaluatePop
	case TOKENS.Do:
		eval = EvaluateDo
	case TOKENS.And:
		eval = EvaluateAnd
	case TOKENS.Or:
		eval = EvaluateOr
	case TOKENS.Fn:
		eval = EvaluateFn
	case TOKENS.Array:
		eval = EvaluateArray
	case TOKENS.Set:
		eval = EvaluateSet
	case TOKENS.Get:
		eval = EvaluateGet
	case TOKENS.If:
		eval = EvaluateIf
	default:
		return nil, errors.New("unknown operator " + exp.Operator.Type)
	}

	return eval(exp)
}

func EvaluateBangEqual(exp *Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Children[0])
	right, rightErr := Evaluate(exp.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	return left != right, nil
}

func EvaluateEqualEqual(exp *Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Children[0])
	right, rightErr := Evaluate(exp.Children[1])

	if leftErr != nil {
		return nil, leftErr
	}

	if rightErr != nil {
		return nil, rightErr
	}

	return left == right, nil
}

func EvaluateGreater(exp *Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Children[0])
	right, rightErr := Evaluate(exp.Children[1])

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

func EvaluateGreaterEqual(exp *Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Children[0])
	right, rightErr := Evaluate(exp.Children[1])

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

func EvaluateLess(exp *Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Children[0])
	right, rightErr := Evaluate(exp.Children[1])

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

func EvaluateLessEqual(exp *Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Children[0])
	right, rightErr := Evaluate(exp.Children[1])

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

func EvaluateMinus(exp *Expression) (Value, error) {
	if len(exp.Children) == 1 {
		return EvaluateMinusUnary(exp)
	} else {
		return EvaluateMinusBinary(exp)
	}
}

func EvaluateMinusBinary(exp *Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Children[0])
	right, rightErr := Evaluate(exp.Children[1])

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

func EvaluateMinusUnary(exp *Expression) (Value, error) {
	right, rightErr := Evaluate(exp.Children[1])

	if rightErr != nil {
		return nil, rightErr
	}

	rightV, ok := right.(float64)

	if !ok {
		return nil, errors.New("right operand is not a number")
	}

	return -rightV, nil
}

func EvaluatePlus(exp *Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Children[0])
	right, rightErr := Evaluate(exp.Children[1])

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

func EvaluateSlash(exp *Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Children[0])
	right, rightErr := Evaluate(exp.Children[1])

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

func EvaluateStar(exp *Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Children[0])
	right, rightErr := Evaluate(exp.Children[1])

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

func EvaluateBang(exp *Expression) (Value, error) {
	right, rightErr := Evaluate(exp.Children[1])

	if rightErr != nil {
		return nil, errors.New("error evaluating inputs")
	}

	rightV, ok := right.(bool)

	if !ok {
		return nil, errors.New("right operand is not a boolean")
	}

	return !rightV, nil
}

func EvaluateTrue(exp *Expression) (Value, error) {
	return true, nil
}

func EvaluateFalse(exp *Expression) (Value, error) {
	return false, nil
}

func EvaluateNil(exp *Expression) (Value, error) {
	return nil, nil
}

func EvaluateNumber(exp *Expression) (Value, error) {
	num, parseErr := strconv.ParseFloat(GetLexemeForToken(exp.Operator), 64)

	if parseErr != nil {
		return nil, errors.New("error parsing number")
	}

	return num, nil
}

func EvaluateString(exp *Expression) (Value, error) {
	return GetLexemeForToken(exp.Operator), nil
}

func EvaluateSet(exp *Expression) (Value, error) {
	dataV, err := Evaluate(exp.Children[0])

	if err != nil {
		return nil, err
	}

	data, ok := dataV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	indexV, err := Evaluate(exp.Children[1])

	index, ok := indexV.(float64)

	if !ok {
		return nil, errors.New("not a number")
	}

	if err != nil {
		return nil, err
	}

	val, err := Evaluate(exp.Children[2])

	if err != nil {
		return nil, err
	}

	data[int(index)] = val

	return data, nil
}

func EvaluateGet(exp *Expression) (Value, error) {
	dataV, err := Evaluate(exp.Children[0])

	if err != nil {
		return nil, err
	}

	data, ok := dataV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	indexV, err := Evaluate(exp.Children[1])

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

func EvaluateFn(exp *Expression) (Value, error) {
	var lambda Lambda = func(arguments ...Value) (Value, error) {
		ClearDefs(exp)

		if len(arguments) != len(exp.Parameters) {
			return nil, errors.New("wrong number of arguments")
		}

		if len(exp.Parameters) > 0 {
			for i, param := range exp.Parameters {
				// TODO, What is the correct way to create new environments for each call?
				err := DefineDef(exp, param, arguments[i])

				if err != nil {
					return nil, err
				}
			}
		}

		var value Value

		for _, child := range exp.Children {
			val, err := Evaluate(child)

			if err != nil {
				return nil, err
			}

			value = val
		}

		return value, nil
	}

	return lambda, nil
}

func EvaluateCall(exp *Expression) (Value, error) {
	fnVal, err := Evaluate(exp.Children[0])

	if err != nil {
		return nil, err
	}

	fn, ok := fnVal.(Lambda)

	if !ok {
		return nil, errors.New("tried to call an expression that didn't evaluate to a Lambda")
	}

	argsExps := exp.Children[1:]

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

	val, err := fn(args...)

	if err != nil {
		return nil, err
	}

	return val, nil
}

func EvaluateDef(exp *Expression) (Value, error) {
	identifier := exp.Children[0].Operator.Lexeme

	var value Value

	for _, input := range exp.Children[1:] {
		val, err := Evaluate(input)

		if err != nil {
			return nil, err
		}

		value = val
	}

	DefineDef(exp.Parent, identifier, value)

	return value, nil
}

func EvaluateLet(exp *Expression) (Value, error) {
	identifier := exp.Children[0].Operator.Lexeme

	var value Value

	for _, input := range exp.Children[1:] {
		val, err := Evaluate(input)

		if err != nil {
			return nil, err
		}

		value = val
	}

	EditDef(exp.Parent, identifier, value)

	return value, nil
}

func EvaluateIdentifier(exp *Expression) (Value, error) {
	if exp.Parent == nil {
		return nil, errors.New("cannot evaluate identifier " + exp.Operator.Lexeme + " with nil parent")
	}
	return ResolveDef(exp.Parent, exp.Operator.Lexeme)
}

func EvaluateDo(exp *Expression) (Value, error) {
	var val Value

	for _, input := range exp.Children {
		v, err := Evaluate(input)

		val = v

		if err != nil {
			return nil, err
		}
	}

	return val, nil
}

func EvaluateIf(exp *Expression) (Value, error) {
	whenExp := exp.Children[0]
	thenExp := exp.Children[1]
	elseExp := exp.Children[2]

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

func EvaluateFor(exp *Expression) (Value, error) {
	arrV, err := Evaluate(exp.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	fnV, err := Evaluate(exp.Children[1])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(Lambda)

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

func EvaluateMap(exp *Expression) (Value, error) {
	arrV, err := Evaluate(exp.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	fnV, err := Evaluate(exp.Children[1])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(Lambda)

	if !ok {
		return nil, errors.New("not a function")
	}

	vals := []Value{}

	for i, v := range arr {
		mapped, err := fn(v, float64(i))

		if err != nil {
			return nil, err
		}

		vals = append(vals, mapped)
	}

	return vals, nil
}

func EvaluatePush(exp *Expression) (Value, error) {
	arrV, err := Evaluate(exp.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	val, err := Evaluate(exp.Children[1])

	if err != nil {
		return nil, err
	}

	arr = append(arr, val)

	return arr, nil
}

func EvaluatePop(exp *Expression) (Value, error) {
	arrV, err := Evaluate(exp.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	return arr[:len(arr)-1], nil
}

func EvaluateFilter(exp *Expression) (Value, error) {
	arrV, err := Evaluate(exp.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	fnV, err := Evaluate(exp.Children[1])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(Lambda)

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

func EvaluateReduce(exp *Expression) (Value, error) {
	arrV, err := Evaluate(exp.Children[0])

	if err != nil {
		return nil, err
	}

	arr, ok := arrV.([]Value)

	if !ok {
		return nil, errors.New("not an array")
	}

	initV, err := Evaluate(exp.Children[1])

	if err != nil {
		return nil, err
	}

	fnV, err := Evaluate(exp.Children[2])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(Lambda)

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

func EvaluateArray(exp *Expression) (Value, error) {
	arr := []Value{}

	for _, input := range exp.Children {
		val, err := Evaluate(input)

		if err != nil {
			return nil, err
		}

		arr = append(arr, val)
	}

	return arr, nil
}

func EvaluateAnd(exp *Expression) (Value, error) {
	if (len(exp.Children) % 2) != 0 {
		return nil, errors.New("and must have even number of inputs")
	}

	var val Value = nil

	for i := 0; i < len(exp.Children); i += 2 {
		conditionVal, err := Evaluate(exp.Children[i])

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

		body, err := Evaluate(exp.Children[i+1])

		if err != nil {
			return nil, err
		}

		val = body
	}

	return val, nil
}

func EvaluateOr(exp *Expression) (Value, error) {
	if (len(exp.Children) % 2) != 0 {
		return nil, errors.New("or must have even number of inputs")
	}

	for i := 0; i < len(exp.Children); i += 2 {
		conditionVal, err := Evaluate(exp.Children[i])

		if err != nil {
			return nil, err
		}

		condition, ok := conditionVal.(bool)

		if !ok {
			return nil, errors.New("condition is not a boolean")
		}

		if condition {
			return Evaluate(exp.Children[i+1])
		}
	}

	return nil, nil
}

func EvaluateLogical(exp *Expression) (Value, error) {
	left, leftErr := Evaluate(exp.Children[0])

	if leftErr != nil {
		return nil, leftErr
	}

	leftV, ok := left.(bool)

	if !ok {
		return nil, errors.New("left operand is not a boolean")
	}

	if exp.Operator.Type == TOKENS.Conjunction {
		if !leftV {
			return false, nil
		}
	} else if exp.Operator.Type == TOKENS.Disjunction {
		if leftV {
			return true, nil
		}
	} else {
		return nil, errors.New("unknown logical operator, && and || are supported")
	}

	right, rightErr := Evaluate(exp.Children[1])

	if rightErr != nil {
		return nil, rightErr
	}

	rightV, ok := right.(bool)

	if !ok {
		return nil, errors.New("right operand is not a boolean")
	}

	return rightV, nil
}
