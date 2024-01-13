package main

import (
	"errors"
)

type Lambda func(args ...Value) (Value, error)

func EvaluateDef(closure *Trajectory) (Value, error) {
	expand(closure)

	nameV, err := evaluate(closure.Children[0])

	if err != nil {
		return nil, err
	}

	name, ok := nameV.(string)

	if !ok {
		return nil, errors.New("not a string")
	}

	fnV, err := evaluate(closure.Children[1])

	if err != nil {
		return nil, err
	}

	fn, ok := fnV.(Lambda)

	if !ok {
		return nil, errors.New("not a function")
	}

	DefineName(closure.Parent, name, fn)

	return fn, nil
}

func EvaluateLambda(closure *Trajectory) (Value, error) {
	var lambda Lambda = func(args ...Value) (Value, error) {

		//
		// EVALUATE FUNCTION BODY
		//

		var value Value

		for _, child := range closure.Expression.Children {
			traj := Traj(&namespace, child)

			val, err := evaluate(&traj)

			if err != nil {
				return nil, err
			}

			value = val
		}

		return value, nil
	}

	return lambda, nil
}

func EvaluateCall(closure *Trajectory) (Value, error) {
	expand(closure)

	lambda, err := ResolveName(closure.Parent, closure.Expression.Operator.Lexeme)

	if err != nil {
		return nil, err
	}

	namespace := Traj(closure, closure.Expression)

	args := []Value{}

	for i, child := range closure.Children {
		arg, err := evaluate(child)

		if err != nil {
			return nil, err
		}

		param := closure.Expression.Parameters[i]

		DefineName(&namespace, param, func(args ...Value) (Value, error) {
			return arg, nil
		})
	}

	return lambda(args)
}
