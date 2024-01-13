package main

import "errors"

func dispatch(trajectory *Trajectory) (Evaluator, error) {
	switch trajectory.Expression.Operator.Type {
	case "ROOT":
		return EvaluateDo, nil
	case TOKENS.BangEqual:
		return EvaluateBangEqual, nil
	case TOKENS.EqualEqual:
		return EvaluateEqualEqual, nil
	case TOKENS.GreaterThan:
		return EvaluateGreater, nil
	case TOKENS.GreaterThanEqual:
		return EvaluateGreaterEqual, nil
	case TOKENS.LessThan:
		return EvaluateLess, nil
	case TOKENS.LessThanEqual:
		return EvaluateLessEqual, nil
	case TOKENS.Conjunction, TOKENS.Disjunction:
		return EvaluateLogical, nil
	case TOKENS.Minus:
		return EvaluateMinus, nil
	case TOKENS.Plus:
		return EvaluatePlus, nil
	case TOKENS.Divide:
		return EvaluateSlash, nil
	case TOKENS.Multiply:
		return EvaluateStar, nil
	case TOKENS.Bang:
		return EvaluateBang, nil
	case TOKENS.True:
		return EvaluateTrue, nil
	case TOKENS.False:
		return EvaluateFalse, nil
	case TOKENS.Nil:
		return EvaluateNil, nil
	case TOKENS.Number:
		return EvaluateNumber, nil
	case TOKENS.String:
		return EvaluateString, nil
	case TOKENS.Identifier:
		return EvaluateIdentifier, nil
	case TOKENS.Def:
		return EvaluateDef, nil
	case TOKENS.Call:
		return EvaluateCall, nil
	case TOKENS.Edit:
		return EvaluateLet, nil
	case TOKENS.Filter:
		return EvaluateFilter, nil
	case TOKENS.For:
		return EvaluateFor, nil
	case TOKENS.Map:
		return EvaluateMap, nil
	case TOKENS.Reduce:
		return EvaluateReduce, nil
	case TOKENS.Push:
		return EvaluatePush, nil
	case TOKENS.Pop:
		return EvaluatePop, nil
	case TOKENS.Do:
		return EvaluateDo, nil
	case TOKENS.And:
		return EvaluateAnd, nil
	case TOKENS.Or:
		return EvaluateOr, nil
	case TOKENS.Fn:
		return EvaluateFn, nil
	case TOKENS.While:
		return EvaluateWhile, nil
	case TOKENS.Array:
		return EvaluateArray, nil
	case TOKENS.Set:
		return EvaluateSet, nil
	case TOKENS.Get:
		return EvaluateGet, nil
	case TOKENS.If:
		return EvaluateIf, nil
	default:
		return nil, errors.New("unknown operator " + trajectory.Expression.Operator.Type)
	}
}
