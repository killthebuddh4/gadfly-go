package main

import "errors"

type Evaluator func(*Trajectory) (Value, error)

func Evaluate(trajectory *Trajectory) (Value, error) {
	switch trajectory.Expression.Variant {
	case VARIANTS.Literal:
		if trajectory.Expression.Operator.Type == TOKENS.Number {
			return EvaluateNumber(trajectory)
		} else if trajectory.Expression.Operator.Type == TOKENS.String {
			return EvaluateString(trajectory)
		} else {
			return nil, errors.New("unknown literal type")
		}
	case VARIANTS.Call:
		lambda, err := ResolveName(trajectory.Parent, trajectory.Expression.Operator.Lexeme)

		if err != nil {
			return nil, err
		}

		return lambda(trajectory)
	default:
		return nil, errors.New("unknown expression variant " + trajectory.Expression.Variant)
	}
}
