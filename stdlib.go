package main

import (
	"errors"
	"fmt"
)

func InitializeStdLib() {
	SetSymbol("print", func(args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, errors.New("print only accepts one argument, a string")
		}

		arg := args[0]

		str, ok := arg.(string)

		if !ok {
			return nil, errors.New("print only accepts strings")
		} else {
			fmt.Println(str)
			return nil, nil
		}
	})
}
