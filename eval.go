package main

import "errors"

type Evaluator func(*Trajectory) (Value, error)

func Evaluate(root *Trajectory) (Value, error) {
	err := InitStdlib(root)

	if err != nil {
		return nil, err
	}

	return evaluate(root)
}

func evaluate(trajectory *Trajectory) (Value, error) {
	switch trajectory.Expression.Variant {
	case VARIANTS.Root:
		return EvaluateDo(trajectory)
	case VARIANTS.Literal:
		switch trajectory.Expression.Operator.Type {
		case TOKENS.Number:
			return EvaluateNumber(trajectory)
		case TOKENS.String:
			return EvaluateString(trajectory)
		case TOKENS.True:
			return EvaluateTrue(trajectory)
		case TOKENS.False:
			return EvaluateFalse(trajectory)
		case TOKENS.Nil:
			return EvaluateNil(trajectory)
		default:
			return nil, errors.New("unknown literal type <" + trajectory.Expression.Operator.Type + ">")
		}
	case VARIANTS.Operator:
		switch trajectory.Expression.Operator.Type {
		case TOKENS.Plus:
			return EvaluatePlus(trajectory)
		case TOKENS.Minus:
			return EvaluateMinus(trajectory)
		case TOKENS.Multiply:
			return EvaluateMultiply(trajectory)
		case TOKENS.Divide:
			return EvaluateDivide(trajectory)
		case TOKENS.EqualEqual:
			return EvaluateEqualEqual(trajectory)
		case TOKENS.BangEqual:
			return EvaluateBangEqual(trajectory)
		case TOKENS.LessThan:
			return EvaluateLessThan(trajectory)
		case TOKENS.LessThanEqual:
			return EvaluateLessThanEqual(trajectory)
		case TOKENS.GreaterThan:
			return EvaluateGreaterThan(trajectory)
		case TOKENS.GreaterThanEqual:
			return EvaluateGreaterThanEqual(trajectory)
		case TOKENS.Bang:
			return EvaluateBang(trajectory)
		case TOKENS.Conjunction:
			return EvaluateConjunction(trajectory)
		case TOKENS.Disjunction:
			return EvaluateDisjunction(trajectory)
		default:
			return nil, errors.New("unknown operator type " + trajectory.Expression.Operator.Type)
		}
	case VARIANTS.Call:
		return EvaluateCall(trajectory)
	default:
		return nil, errors.New("unknown expression variant <" + trajectory.Expression.Variant + "> for token <" + trajectory.Expression.Operator.Lexeme + ">")
	}
}
