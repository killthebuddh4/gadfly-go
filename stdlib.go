package main

import (
	"errors"
	"fmt"
)

func InitializeStdLib() {
	DefSymbol("print", func(args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, errors.New("print only accepts one argument, a string")
		}

		arg := args[0]

		str, strOk := arg.(string)
		num, numOk := arg.(float64)

		if strOk {
			fmt.Println(str)
			return nil, nil
		} else if numOk {
			fmt.Println(num)
			return nil, nil
		} else {
			return nil, errors.New("print only accepts strings and numbers")
		}
	})
}
