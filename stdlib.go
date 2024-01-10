package main

import (
	"errors"
	"fmt"
	"reflect"
)

func InitializeStdLib(root *Trajectory) {
	var lambda Lambda = func(args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, errors.New("print only accepts one argument, a string")
		}

		arg := args[0]

		str, strOk := arg.(string)
		float, floatOk := arg.(float64)
		i, intOk := arg.(int)
		tf, tfOk := arg.(bool)
		slice, sliceOk := arg.([]Value)

		if strOk {
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

	DefineName(root, "print", lambda)
}
