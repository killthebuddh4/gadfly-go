package main

import (
	"errors"
	"fmt"
	"reflect"
)

func EvaluatePrint(scope *Trajectory) (Value, error) {
	expand(scope)

	arg, err := evaluate(scope.Children[0])

	if err != nil {
		return nil, err
	}

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
