package main

import (
	"errors"
	"fmt"
)

type Lambda func(scope *Trajectory, args ...Value) (Value, error)

func EvaluateDef(scope *Trajectory, args ...Value) (Value, error) {
	fmt.Println("Evaluating def", args)

	name, ok := args[0].(string)

	if !ok {
		return nil, errors.New("trying to define a name that is is not a string")
	}

	fn, ok := args[1].(Lambda)

	if !ok {
		return nil, errors.New("trying to define a value that's not a function")
	}

	DefineName(scope, name, fn)

	return name, nil
}
