package main

import "errors"

func EvaluateFn(trajectory *Trajectory) (Value, error) {
	var lambda Lambda = func(arguments ...Value) (Value, error) {
		if len(arguments) != len(trajectory.Expression.Parameters) {
			return nil, errors.New("wrong number of arguments")
		}

		scope := Traj(trajectory, trajectory.Expression)

		for i, param := range trajectory.Expression.Parameters {
			DefineName(&scope, param, arguments[i])
		}

		var value Value

		for _, child := range trajectory.Expression.Children {
			traj := Traj(&scope, child)

			val, err := Evaluate(&traj)

			if err != nil {
				return nil, err
			}

			value = val
		}

		return value, nil
	}

	return lambda, nil
}

func EvaluateCall(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	fnVal, err := Evaluate(trajectory.Children[0])

	if err != nil {
		return nil, err
	}

	fn, ok := fnVal.(Lambda)

	if !ok {
		return nil, errors.New("tried to call an expression that didn't evaluate to a Lambda")
	}

	argsT := trajectory.Children[1:]

	args := []Value{}

	if len(argsT) > 0 {
		for _, traj := range argsT {
			arg, err := Evaluate(traj)

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

func EvaluateDef(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	// TODO
	identifier := trajectory.Children[0].Expression.Operator.Lexeme

	var value Value

	for _, input := range trajectory.Children[1:] {
		val, err := Evaluate(input)

		if err != nil {
			return nil, err
		}

		value = val
	}

	DefineName(trajectory.Parent, identifier, value)

	return value, nil
}

func EvaluateLet(trajectory *Trajectory) (Value, error) {
	expand(trajectory)
	// TODO
	identifier := trajectory.Children[0].Expression.Operator.Lexeme

	var value Value

	for _, input := range trajectory.Children[1:] {
		val, err := Evaluate(input)

		if err != nil {
			return nil, err
		}

		value = val
	}

	EditName(trajectory.Parent, identifier, value)

	return value, nil
}

func EvaluateIdentifier(trajectory *Trajectory) (Value, error) {
	if trajectory.Parent == nil {
		return nil, errors.New("cannot evaluate identifier " + trajectory.Expression.Operator.Lexeme + " with nil parent")
	}
	return ResolveName(trajectory.Parent, trajectory.Expression.Operator.Lexeme)
}
