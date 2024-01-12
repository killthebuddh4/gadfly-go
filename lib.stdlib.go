package main

import (
	"errors"
	"fmt"
	"reflect"
)

func InitializeStdLib(root *Trajectory) {
	DefineName(root, TOKENS.Def, EvaluateDef)
	DefineName(root, TOKENS.BangEqual, EvaluateBangEqual)
	DefineName(root, TOKENS.EqualEqual, EvaluateEqualEqual)
	DefineName(root, TOKENS.LessThan, EvaluateLessThan)
	DefineName(root, TOKENS.LessThanEqual, EvaluateLessThanEqual)
	DefineName(root, TOKENS.GreaterThan, EvaluateGreaterThan)
	DefineName(root, TOKENS.GreaterThanEqual, EvaluateGreaterThanEqual)
	DefineName(root, TOKENS.Plus, EvaluatePlus)
	DefineName(root, TOKENS.Minus, EvaluateMinus)
	DefineName(root, TOKENS.Multiply, EvaluateMultiply)
	DefineName(root, TOKENS.Divide, EvaluateDivide)
	DefineName(root, TOKENS.Bang, EvaluateBang)
	DefineName(root, TOKENS.True, EvaluateTrue)
	DefineName(root, TOKENS.False, EvaluateFalse)
	DefineName(root, TOKENS.Nil, EvaluateNil)

	var print Lambda = func(scope *Trajectory, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, errors.New("print only accepts one argument, a string")
		}

		arg := args[0]

		str, strOk := arg.(string)
		float, floatOk := arg.(float64)
		i, intOk := arg.(int)
		tf, tfOk := arg.(bool)
		slice, sliceOk := arg.([]Value)

		if arg == nil {
			fmt.Println("nil")
		} else if strOk {
			fmt.Println(str)
		} else if floatOk {
			fmt.Println(float)
		} else if intOk {
			fmt.Println(i)
		} else if tfOk {
			fmt.Println(tf)
		} else if sliceOk {
			fmt.Println("[")
			for _, v := range slice {
				fmt.Print("    ")
				fmt.Println(v)
			}
			fmt.Println("]")
		} else {
			return nil, errors.New("print only accepts booleans and strings and numbers, got " + reflect.TypeOf(arg).String())
		}

		return nil, nil
	}

	DefineName(root, "print", print)

	// var chars Lambda = func(args ...Value) (Value, error) {
	// 	if len(args) != 1 {
	// 		return nil, errors.New("chars only accepts one argument, a string")
	// 	}

	// 	arg := args[0]

	// 	str, strOk := arg.(string)

	// 	if !strOk {
	// 		return nil, errors.New("chars only accepts strings")
	// 	}

	// 	result := []Value{}

	// 	for _, c := range str {
	// 		result = append(result, string(c))
	// 	}

	// 	return result, nil
	// }

	// DefineName(root, "chars", chars)
}
