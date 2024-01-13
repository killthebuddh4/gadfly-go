package main

import "errors"

func EvaluateIf(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	whenExp := trajectory.Children[0]
	thenExp := trajectory.Children[1]
	elseExp := trajectory.Children[2]

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

func EvaluateAnd(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	if (len(trajectory.Children) % 2) != 0 {
		return nil, errors.New("and must have even number of inputs")
	}

	var val Value = nil

	for i := 0; i < len(trajectory.Children); i += 2 {
		conditionVal, err := Evaluate(trajectory.Children[i])

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

		body, err := Evaluate(trajectory.Children[i+1])

		if err != nil {
			return nil, err
		}

		val = body
	}

	return val, nil
}

func EvaluateOr(trajectory *Trajectory) (Value, error) {
	expand(trajectory)

	if (len(trajectory.Children) % 2) != 0 {
		return nil, errors.New("or must have even number of inputs")
	}

	for i := 0; i < len(trajectory.Children); i += 2 {
		conditionVal, err := Evaluate(trajectory.Children[i])

		if err != nil {
			return nil, err
		}

		condition, ok := conditionVal.(bool)

		if !ok {
			return nil, errors.New("condition is not a boolean")
		}

		if condition {
			return Evaluate(trajectory.Children[i+1])
		}
	}

	return nil, nil
}
